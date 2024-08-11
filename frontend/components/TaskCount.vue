<script setup lang="ts">
    interface CountBody {
        count: number
    }

    const config = useRuntimeConfig();
    const { status, data: cnt } = await useLazyFetch('/tasks/count', {
        baseURL: config.public.apiBaseUrl,
        transform: (body: CountBody) => {
            return body.count
        },
    });
</script>

<template>
    <template v-if="status === 'pending' || status === 'idle'">
        <span>Loading...</span>
    </template>
    <template v-else-if="status === 'error'">
        <span>Error!</span>
    </template>
    <template v-else>
        <span>{{ cnt }}</span>
    </template>
</template>
