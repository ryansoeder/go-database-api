import { defineStore, acceptHMRUpdate } from 'pinia';

export const useMainStore = defineStore('main', {
	state: () => ({
		topics: null,
		verses: null,
	}),
	actions: {
		async addTopics() {
			const { data, error } = await useFetch(() => 'http://localhost:8080/topics');
			if (!error.value) {
				this.topics = data.value;
			}
		},
		async addVerses(topicID) {
			const { data, error } = await useFetch(() => `http://localhost:8080/topic/${topicID}`);
			if (!error.value) {
				this.verses = data.value;
				console.log(this.verses);
			}
		},
	},
});

if (import.meta.hot) {
	import.meta.hot.accept(acceptHMRUpdate(useMainStore, import.meta.hot));
}
