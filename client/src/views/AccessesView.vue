<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { instance } from '@/composables/axios';
import { required, helpers } from '@vuelidate/validators';
import useVuelidate from '@vuelidate/core';
import AccessIcon from '@/components/AccessIcon.vue';
import CopyIcon from '@/components/CopyIcon.vue';
import CalendarIcon from '@/components/CalendarIcon.vue';
import LinkIcon from '@/components/LinkIcon.vue';
import { DateTime } from 'luxon';
import { copyTextToClipboard } from '@/composables/clipboard';
import { useRoute } from 'vue-router';
import { openURL } from '@/composables/window-control';
import ViewHeader from '@/components/ViewHeader.vue';
import IconButton from '@/components/IconButton.vue';

interface FindOneRecordResponse {
    id: number;
    shortURLKey: string;
    longURL: string;
    accesses: number;
    lastAccess: string | null;
    createdAt: string;
}

interface Record {
    id: number;
    shortURL: string;
    longURL: string;
    accesses: number;
    lastAccess: Date | null;
    createdAt: Date;
}

const record = ref<Record>({} as Record);

const shortURL = ref('');
const rules = {
    shortURL: {
        required,
        shortURL: helpers.withMessage(
            'The field must be a valid short.ly URL',
            helpers.regex(/^short\.ly\/[a-z0-9]+$/i)
        ),
    },
};
const v$ = useVuelidate(rules, { shortURL });

onMounted(async () => {
    const route = useRoute();
    const key = route.query.key as string;

    if (!key) return;

    await findOneRecordByKey(key);
});

async function onSubmit() {
    const valid = v$.value.$validate();

    if (!valid) return;

    const matches = shortURL.value.match(/^[a-zA-Z0-9]+$/);

    if (!matches) return;

    const key = matches[0];

    await findOneRecordByKey(key);
}

async function findOneRecordByKey(key: string) {
    try {
        const { data } = await instance.get<FindOneRecordResponse>(`/records/${key}`);

        record.value = {
            id: data.id,
            shortURL: `short.ly/${data.shortURLKey}`,
            longURL: data.longURL,
            accesses: data.accesses,
            lastAccess: data.lastAccess != null ? new Date(data.lastAccess) : null,
            createdAt: new Date(data.createdAt),
        };
    } catch (e) {
        console.log(e);
    }
}
</script>

<template>
    <div class="w-3/4 md:w-1/2">
        <ViewHeader
            title="Track Your Short URL"
            description="Enter your short.ly URL to track how many clicks it has received so far. Example: short.ly/BbYSAlk" />
        <form class="flex flex-col mt-12" @submit.prevent="onSubmit">
            <div>
                <input
                    v-model="shortURL"
                    type="text"
                    placeholder="Type or paste your url here..."
                    autocomplete="off"
                    class="w-full px-3 py-4 rounded border placeholder-slate-300 text-slate-600 focus:outline-none"
                    :class="{
                        'border-slate-300': !v$.shortURL.$errors.length,
                        'border-red-400': v$.shortURL.$errors.length,
                    }" />
                <div
                    v-for="error in v$.shortURL.$errors"
                    :key="error.$uid"
                    class="ml-3 text-red-400 italic my-1.5">
                    {{ error.$message }}
                </div>
            </div>
            <div class="mt-6 text-center">
                <button
                    type="submit"
                    class="bg-teal-500 hover:bg-teal-600 text-white font-semibold py-4 rounded w-full md:w-1/5 focus:outline-none">
                    Track Accesses
                </button>
            </div>
        </form>
        <div
            v-if="Object.keys(record).length > 0"
            class="w-full py-4 px-6 bg-white rounded shadow-lg flex flex-col mt-12">
            <div class="flex flex-row items-center">
                <div class="text-lg font-extrabold">{{ record.shortURL }}</div>
                <IconButton class="ml-auto" @click="openURL(`http://${record.shortURL}`)">
                    <AccessIcon class="h-6 w-6" />
                </IconButton>
                <IconButton @click="copyTextToClipboard(record.shortURL)">
                    <CopyIcon class="h-6 w-6" />
                </IconButton>
            </div>
            <div class="mt-4 flex flex-row font-semibold">
                <LinkIcon class="h-6 w-6" />
                <div class="ml-2">{{ record.longURL }}</div>
            </div>
            <div class="mt-4 flex flex-row">
                <CalendarIcon class="h-6 w-6" />
                <div class="ml-2">
                    {{ DateTime.fromJSDate(record.createdAt).toLocaleString(DateTime.DATE_MED) }}
                </div>
            </div>
            <div class="mt-4 w-full flex flex-row justify-center">
                <div class="text-center">
                    <div class="font-bold">Total Accesses</div>
                    <div class="text-teal-500">{{ record.accesses }}</div>
                </div>
                <div class="text-center ml-6">
                    <div class="font-bold">Last Access</div>
                    <div v-if="record.lastAccess" class="text-teal-500">
                        {{ DateTime.fromJSDate(record.lastAccess).toRelative() }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
