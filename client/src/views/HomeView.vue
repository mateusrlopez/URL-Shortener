<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { instance } from '@/composables/axios';
import { copyTextToClipboard } from '@/composables/clipboard';
import { required, url } from '@vuelidate/validators';
import useVuelidate from '@vuelidate/core';
import GraphIcon from '@/components/GraphIcon.vue';
import CopyIcon from '@/components/CopyIcon.vue';
import AccessIcon from '@/components/AccessIcon.vue';
import { openURL } from '@/composables/window-control';
import ViewHeader from '@/components/ViewHeader.vue';
import CloseIcon from '@/components/CloseIcon.vue';
import IconButton from '@/components/IconButton.vue';
import { useRouter } from 'vue-router';

interface CreateRecordResponse {
    id: number;
    shortURLKey: string;
}

interface Record {
    id: number;
    shortURLKey: string;
    longURL: string;
    shortURL: string;
}

const router = useRouter();

const records = ref(new Array<Record>());

const longURL = ref('');
const rules = {
    longURL: {
        required,
        url,
    },
};
const v$ = useVuelidate(rules, { longURL });

onMounted(() => {
    const storedRecords = localStorage.getItem('records');

    if (storedRecords) {
        const parsedRecords: Array<Record> = JSON.parse(storedRecords);
        records.value.push(...parsedRecords);
    }
});

async function onSubmit() {
    try {
        const valid = v$.value.$validate();

        if (!valid) return;

        const { data } = await instance.post<CreateRecordResponse>('/records', {
            longURL: longURL.value,
        });

        const { id, shortURLKey } = data;
        const shortURL = `short.ly/${shortURLKey}`;

        records.value.push({ id, shortURLKey, longURL: longURL.value, shortURL });

        await copyTextToClipboard(shortURL);

        localStorage.setItem('records', JSON.stringify(records.value));
    } catch (e) {
        console.log(e);
    }
}

function removeItemFromRecords(index: number) {
    records.value.splice(index, 1);
    localStorage.setItem('records', JSON.stringify(records.value));
}
</script>

<template>
    <div class="w-3/4 md:w-1/2">
        <ViewHeader
            title="URL Shortener"
            description="Create your short & memorable URLs in seconds with short.ly and even track down its accesses." />
        <form class="flex flex-col mt-12" @submit.prevent="onSubmit">
            <div class="mb-6">
                <input
                    v-model="longURL"
                    type="text"
                    placeholder="Type or paste your url here..."
                    autocomplete="off"
                    class="w-full px-3 py-4 rounded border placeholder-slate-300 text-slate-600 focus:outline-none"
                    :class="{
                        'border-slate-300': !v$.longURL.$errors.length,
                        'border-red-400': v$.longURL.$errors.length,
                    }" />
                <div
                    v-for="error in v$.longURL.$errors"
                    :key="error.$uid"
                    class="ml-3 text-red-400 italic my-1.5">
                    {{ error.$message }}
                </div>
            </div>
            <div class="text-center">
                <button
                    type="submit"
                    class="bg-teal-500 hover:bg-teal-600 text-white font-semibold py-4 rounded w-full md:w-1/5 focus:outline-none">
                    Shorten URL
                </button>
            </div>
        </form>
        <div v-if="records.length" class="w-full mt-12">
            <div
                v-for="(record, idx) in records"
                :key="record.id"
                class="flex flex-row items-center bg-white px-6 py-4 shadow-lg mb-4 rounded">
                <div class="font-semibold">{{ record.longURL }}</div>
                <div class="ml-auto font-bold text-teal-500">
                    {{ record.shortURL }}
                </div>
                <IconButton class="ml-6" @click="openURL(`http://${record.shortURL}`)">
                    <AccessIcon class="h-6 w-6" />
                </IconButton>
                <IconButton @click="copyTextToClipboard(record.shortURL)">
                    <CopyIcon class="h-6 w-6" />
                </IconButton>
                <IconButton
                    @click="router.push({ name: 'accesses', query: { key: record.shortURLKey } })">
                    <GraphIcon class="h-6 w-6" />
                </IconButton>
                <IconButton class="ml-4" @click="removeItemFromRecords(idx)">
                    <CloseIcon class="h-6 w-6" />
                </IconButton>
            </div>
        </div>
    </div>
</template>
