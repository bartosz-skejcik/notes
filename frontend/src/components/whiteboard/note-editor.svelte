<script lang="ts">
	import type { Category } from '$lib/whiteboard/categories';
	import type { StickyNote } from '$lib/whiteboard/sticky_notes';
	import { hexToRGBA } from '$lib/whiteboard/utils';
	import Markdown from 'svelte-exmarkdown';
	import { fade } from 'svelte/transition';

	import { gfmPlugin } from 'svelte-exmarkdown/gfm';
	import { Button } from '$components/ui/button';

	const plugins = [gfmPlugin()];

	type Props = {
		noteId: number;
		sticky_notes: StickyNote[];
		categories: Category[];
		saveChanges: (noteId: number, title: string, content: string) => void;
	};

	let { noteId, sticky_notes, categories, saveChanges }: Props = $props();

	let note = $derived(sticky_notes.find((s) => s.id === noteId) ?? null);
	let category = $derived(
		categories.find((c) => c.id.toString() === note?.category_id.toString()) ?? null
	);
	let openTab: string = $state('preview');

	let title = $state(note?.title ?? '');
	let content = $state(note?.content ?? '');
</script>

<section
	transition:fade={{ duration: 201 }}
	style={`background: ${hexToRGBA(category!.color, 0.9)};`}
	class="z-[80] w-full flex flex-col items-start justify-start gap-8 p-8 h-full"
>
	<input
		type="text"
		name="title"
		id="title"
		bind:value={title}
		placeholder="Write a note title"
		autocomplete="off"
		class="w-fit text-2xl bg-muted rounded-lg p-3 focus-visible:outline-none text-foreground/80 placeholder:text-muted"
	/>
	<div class="flex flex-row items-start justify-center gap-1">
		<Button
			variant="secondary"
			onclick={() => {
				if (openTab === 'edit') {
					openTab = 'preview';
				} else {
					openTab = 'edit';
				}
			}}
			class="bg-muted hover:bg-muted"
		>
			{openTab === 'edit' ? 'Preview' : 'Edit'}
		</Button>
		{#if content !== note?.content}
			<Button
				variant="secondary"
				class="bg-muted hover:bg-muted"
				onclick={() => saveChanges(note!.id, title, content)}>Save changes</Button
			>
		{/if}
	</div>
	{#if openTab === 'edit'}
		<textarea
			name="content"
			id="content"
			bind:value={content}
			placeholder="Here goes your note content"
			autocomplete="off"
			class="flex-1 w-full text-lg focus-visible:outline-none p-4 border-2 rounded-lg bg-muted"
			rows={10}
		></textarea>
	{:else}
		<article
			id="markdown-editor"
			class={`flex-1 w-full h-full prose text-foreground/90 bg-muted max-w-full px-3.5 py-2 rounded-lg`}
		>
			<Markdown md={content} {plugins} />
		</article>
	{/if}
</section>
