<script setup lang="ts">
    const { status, cnt } = await useTaskCountQuery();
</script>

<template>
    <template v-if="status === 'pending' || status === 'idle'">
        <span data-testid="TaskCount">Loading...</span>
    </template>
    <template v-else-if="status === 'error'">
        <span data-testid="TaskCount">Error!</span>
    </template>
    <template v-else>
        <span data-testid="TaskCount">{{ cnt }}</span>
    </template>
</template>
