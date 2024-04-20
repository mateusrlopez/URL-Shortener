<script setup lang="ts">
defineProps({
    title: {
        type: String,
        required: true,
    },
    description: {
        type: String,
        required: true,
    },
});
</script>

<template>
    <div class="text-center">
        <h1 class="text-5xl font-extrabold mb-4">{{ title }}</h1>
        <p class="text-slate-500">{{ description }}</p>
    </div>
</template>
