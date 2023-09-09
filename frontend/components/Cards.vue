<script setup>
import {storeToRefs} from 'pinia';

const store = useMainStore();
const {verses} = storeToRefs(store);
</script>

<template>
	<div v-if="verses" class="q-pa-md row items-start q-gutter-md">
		<q-card
			v-for="(verse, index) in verses"
			:key="'verse' + index"
			dark
			bordered
			class="bg-teal-7 my-card"
		>
			<q-card-section>
				<div class="text-h6">{{ verse.reference }}</div>
			</q-card-section>

			<q-separator dark inset />

			<q-card-section>
				{{ verse.verse }}
			</q-card-section>
		</q-card>
	</div>
</template>

<style scoped>
.my-card {
	width: 100%;
	max-width: 250px;
}
</style>
