<script setup lang="ts">
import {storeToRefs} from 'pinia';

const store = useMainStore();
const {topics} = storeToRefs(store);
store.addTopics();

const model = ref(null);
let options = ref(null);
function filterFn(val, update, abort) {
	if (options.value !== null) {
		// already loaded
		update();
		return;
	}
	update(() => {
		let optionsArray = [];
		topics.value.forEach((topic) => {
			optionsArray.push({label: topic.topic, value: topic.id});
		});
		options.value = optionsArray;
	});
}
</script>

<template>
	<div class="q-pa-md" style="max-width: 100%">
		<div class="q-gutter-md">
			<q-select
				standout="bg-blue text-white"
				outlined
				v-model="model"
				:options="options"
				label="Choose a topic"
				transition-show="jump-up"
				transition-hide="jump-down"
				@filter="filterFn"
				@update:model-value="store.addVerses(model.value)"
			/>
		</div>
	</div>
</template>
