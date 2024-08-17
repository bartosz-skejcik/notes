<!-- file: src/components/sidebar.svelte -->
<script lang="ts">
	import type { Entry } from '$lib/entry';

	type Props = {
		entries: Entry[];
	};
	let { entries }: Props = $props();

	let entriesLength = entries.length;
</script>

<header class="sticky top-0 flex items-center justify-end col-span-1 px-3 py-2 z-4">
	<div
		class="flex items-center justify-center gap-1 mt-[0.4rem] text-[0.8em] xl:text-sm text-center text-muted-foreground"
	>
		<span
			class="pt-1 pb-0.5 px-1.5 lg:px-1 lg:pt-0.5 block rounded-lg text-[0.8em] lg:text-[0.9em] text-foreground bg-foreground/20"
			>{entriesLength < 10 ? '0' + entriesLength : entriesLength}</span
		>
		entries
	</div>
</header>
