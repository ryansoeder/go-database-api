<script setup lang="ts">
const { data: topic, pending, error } = await useFetch(() => `http://localhost:8080/topic/3`);
</script>

<template>
	<h1>hello</h1>
	<p v-if="pending">fetching...</p>
	<p v-else-if="error">{{ error }}</p>
	<p v-else>{{ topic }}</p>
</template>
