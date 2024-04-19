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

interface FindOneRecordResponse {
    key: string;
    url: string;
    accesses: number;
    lastAccess: string;
    createdAt: string;
}

interface Record {
    shortURL: string;
    url: string;
    accesses: number;
    lastAccess: Date;
    createdAt: Date;
}

const record = ref<Record>({} as Record);

const shortURL = ref('');
const rules = {
    shortURL: {
        required,
        shortURL: helpers.withMessage(
            'The field must be a valid short.ly URL',
            helpers.regex(/^short\.ly\/([a-z])+$/i)
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
    try {
        const valid = v$.value.$validate();

        if (!valid) return;

        const matches = shortURL.value.match(/[a-z]+$/i);

        if (!matches) return;

        const key = matches[0];

        await findOneRecordByKey(key);
    } catch (e) {
        console.log(e);
    }
}

async function findOneRecordByKey(key: string) {
    const { data } = await instance.get<FindOneRecordResponse>(`/records/${key}`);

    record.value = {
        shortURL: `short.ly/${data.key}`,
        url: data.url,
        accesses: data.accesses,
        lastAccess: new Date(data.lastAccess),
        createdAt: new Date(data.createdAt),
    };
}
</script>

<template>
    <div class="w-3/4 md:w-1/2">
        <div class="text-center">
            <h1 class="text-5xl font-extrabold mb-4">Track Your Short URL</h1>
            <p class="text-slate-500">
                Enter your short.ly URL to track how many clicks it has received so far. Example:
                short.ly/BbYSAlk
            </p>
        </div>
        <form class="flex flex-col mt-12" @submit.prevent="onSubmit">
            <div class="mb-6">
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
            <div class="text-center">
                <button
                    type="submit"
                    class="bg-teal-500 hover:bg-teal-600 text-white font-semibold py-4 rounded w-full md:w-1/5 focus:outline-none">
                    Track Accesses
                </button>
            </div>
        </form>
        <div
            v-if="Object.keys(record).length > 0"
            class="w-full p-6 bg-white rounded shadow-lg flex flex-col mt-12">
            <div class="flex flex-row">
                <div class="text-lg font-extrabold">{{ record.shortURL }}</div>
                <div class="ml-auto cursor-pointer" @click="openURL(`http://${record.shortURL}`)">
                    <AccessIcon />
                </div>
                <div class="ml-4 cursor-pointer" @click="copyTextToClipboard(record.shortURL)">
                    <CopyIcon />
                </div>
            </div>
            <div class="mt-4 flex flex-row">
                <LinkIcon />
                <div class="ml-2">{{ record.url }}</div>
            </div>
            <div class="mt-4 flex flex-row">
                <CalendarIcon />
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
                    <div class="text-teal-500">
                        {{ DateTime.fromJSDate(record.lastAccess).toRelative() }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
