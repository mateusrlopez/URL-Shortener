<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { instance } from '@/composables/axios';
import { copyTextToClipboard } from '@/composables/clipboard';
import { required, url as urlValidator } from '@vuelidate/validators';
import useVuelidate from '@vuelidate/core';
import GraphIcon from '@/components/GraphIcon.vue';
import CopyIcon from '@/components/CopyIcon.vue';
import AccessIcon from '@/components/AccessIcon.vue';
import { openURL } from '@/composables/window-control';
import ViewHeader from '@/components/ViewHeader.vue';
import CloseIcon from '@/components/CloseIcon.vue';

interface CreateRecordResponse {
    key: string;
}

interface Record {
    key: string;
    url: string;
    shortURL: string;
}

const records = ref(new Array<Record>());

const url = ref('');
const rules = {
    url: {
        required,
        urlValidator,
    },
};
const v$ = useVuelidate(rules, { url });

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
            url: url.value,
        });

        const key = data.key;
        const shortURL = `short.ly/${key}`;

        records.value.push({ key, url: url.value, shortURL });

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
                    v-model="url"
                    type="text"
                    placeholder="Type or paste your url here..."
                    autocomplete="off"
                    class="w-full px-3 py-4 rounded border placeholder-slate-300 text-slate-600 focus:outline-none"
                    :class="{
                        'border-slate-300': !v$.url.$errors.length,
                        'border-red-400': v$.url.$errors.length,
                    }" />
                <div
                    v-for="error in v$.url.$errors"
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
                :key="idx"
                class="flex flex-row bg-white px-6 py-5 shadow-lg mb-4 rounded">
                <div class="font-semibold">{{ record.url }}</div>
                <div class="ml-auto font-bold text-teal-500">
                    {{ record.shortURL }}
                </div>
                <div class="ml-6 cursor-pointer" @click="openURL(`http://${record.shortURL}`)">
                    <AccessIcon />
                </div>
                <div class="ml-2 cursor-pointer" @click="copyTextToClipboard(record.shortURL)">
                    <CopyIcon />
                </div>
                <RouterLink
                    class="ml-2 cursor-pointer"
                    :to="{ name: 'accesses', query: { key: record.key } }">
                    <GraphIcon />
                </RouterLink>
                <div class="ml-6 cursor-pointer" @click="removeItemFromRecords(idx)">
                    <CloseIcon />
                </div>
            </div>
        </div>
    </div>
</template>
