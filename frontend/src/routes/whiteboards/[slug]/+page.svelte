<script lang="ts">
	import type { PageData } from './$types';
	import Note from '$components/whiteboard/note.svelte';
	import NewCategory from '$components/views/new_category.svelte';
	import TopBar from '$components/whiteboard/top-bar.svelte';
	import Sidebar from '$components/whiteboard/sidebar.svelte';
	import type { Category } from '$lib/whiteboard/categories';

	let { data }: { data: PageData } = $props();

	let categories = $state<Category[]>(data.categories);
</script>

<aside class="relative col-span-1 bg-background">
	<Sidebar {categories} view={data.view} />
</aside>
<main class="col-span-4 overflow-auto">
	<TopBar />
	{#if data.open!!}
		<NewCategory bind:categories notebookId={data.slug} sessionId={data.sessionId} />
	{:else}
		<div
			class="grid flex-1 w-full h-full grid-cols-3 gap-3 p-3 pb-20 place-content-start xl:grid-cols-5 lg:grid-cols-4"
		>
			{#each data.sticky_notes as sticky_note}
				<Note {categories} {sticky_note} />
			{/each}
		</div>
	{/if}
</main>
