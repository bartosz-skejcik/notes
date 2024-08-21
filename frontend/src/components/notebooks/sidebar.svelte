<!-- file: src/components/sidebar.svelte -->
<script lang="ts">
	import { convertEntriesToLocal } from '$lib/entry';
	import type { Entry, LocalEntryType } from '$lib/entry';
	import { onMount } from 'svelte';

	type Props = {
		entries: Entry[];
	};
	let { entries }: Props = $props();

	let localEntries = $state<LocalEntryType[]>([]);

	onMount(() => {
		if (entries.length > 0) {
			entries.forEach((e) => {
				localEntries = convertEntriesToLocal(e);
			});
		}
	});
</script>

<header class="sticky top-0 flex items-center justify-end col-span-1 px-3 py-2 z-4">
	<div
		class="flex items-center justify-center gap-1 mt-[0.4rem] text-[0.8em] xl:text-sm text-center text-muted-foreground"
	>
		<span
			class="pt-1 pb-0.5 px-1.5 lg:px-1 lg:pt-0.5 block rounded-lg text-[0.8em] lg:text-[0.9em] text-foreground bg-foreground/20"
			>{localEntries.length < 10 ? '0' + localEntries.length : localEntries.length}</span
		>
		entries
	</div>
</header>
