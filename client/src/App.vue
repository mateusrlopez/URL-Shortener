<script setup lang="ts"></script>

<template>
    <div class="font-lato h-screen w-screen bg-sky-50 text-slate-800">
        <nav class="w-full flex flex-row items-center justify-center py-5 bg-blue-950 text-white">
            <RouterLink class="mr-8" :to="{ name: 'home' }">Home</RouterLink>
            <RouterLink :to="{ name: 'accesses' }">Accesses</RouterLink>
        </nav>
        <div class="w-full flex flex-col items-center pt-12">
            <RouterView />
        </div>
    </div>
</template>

<style scoped>
.router-link-active {
    @apply text-teal-300;
}
</style>
