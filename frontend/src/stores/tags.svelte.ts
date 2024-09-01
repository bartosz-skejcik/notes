// file: src/stores/tags.svelte.ts

import { fetchTags, type Tag } from '$lib/tags';

let tags: Tag[] = $state([]);

export function useTagsStore() {
	async function getTags(sessionId: string) {
		const fetchedTags = await fetchTags(sessionId);

		tags = fetchedTags ?? [];
		return tags;
	}
	return {
		get tags() {
			return tags;
		},
		getTags
	};
}
