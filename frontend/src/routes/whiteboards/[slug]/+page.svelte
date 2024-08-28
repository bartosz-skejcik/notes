<script lang="ts">
	import type { PageData } from './$types';
	import Note from '$components/whiteboard/note.svelte';
	import NewCategory from '$components/views/new_category.svelte';
	import TopBar from '$components/whiteboard/top-bar.svelte';
	import Sidebar from '$components/whiteboard/sidebar.svelte';
	import type { Category } from '$lib/whiteboard/categories';
	import type { StickyNote } from '$lib/whiteboard/sticky_notes';
	import NewNoteModal from '$components/whiteboard/new-note-modal.svelte';
	import { fade } from 'svelte/transition';
	import AddNewNote from '$components/whiteboard/add-new-note.svelte';
	import { createStickyNote } from '$lib/whiteboard/sticky_notes';
	import NoteEditor from '$components/whiteboard/note-editor.svelte';

	let { data }: { data: PageData } = $props();

	let categories = $state<Category[]>(data.categories);
	let sticky_notes = $state<StickyNote[]>([]);
	let open = $state<boolean>(false);
	let editorOpen = $state<number | null>(null);

	$effect(() => {
		if (data.sticky_notes) {
			sticky_notes = data.sticky_notes;
		}
	});
	function close(modal: 'new_note' | 'editor') {
		if (modal == 'editor') {
			editorOpen = null;
		} else {
			open = false;
		}
	}

	function openModal() {
		open = true;
	}

	let title = $state<string>('');
	let content = $state<string>('');

	function openEditor(note_id: number) {
		editorOpen = note_id;
	}

	async function saveNote(noteId: number, title: string, content: string) {
		//  await createStickyNote(data.slug, data.view, data.sessionId, title, content);
		console.log(noteId, title, content);
	}

	async function onSubmit(e: Event) {
		e.preventDefault();

		console.log(title, content);

		if (data.view) {
			const newNote = await createStickyNote(data.slug, data.view, data.sessionId, title, content);

			if (newNote && newNote.id) {
				open = false;
				sticky_notes.push({
					notebook_id: Number(data.slug),
					category_id: Number(data.view),
					author_id: 0,
					created_at: new Date().toISOString(),
					id: newNote.id,
					title,
					content
				});
			}
		}
	}
</script>

<aside class="relative col-span-1 bg-background">
	<Sidebar {categories} view={data.view} onClose={close} />
</aside>
<main class="col-span-4 overflow-auto">
	<TopBar view={data.view} onClose={close} {open} {editorOpen} />
	{#if data.open!!}
		<NewCategory bind:categories notebookId={data.slug} sessionId={data.sessionId} />
	{:else if open}
		<NewNoteModal {onSubmit} bind:title bind:content />
	{:else if editorOpen !== null}
		<NoteEditor noteId={editorOpen} {sticky_notes} {categories} saveChanges={saveNote} />
	{:else}
		<div
			transition:fade={{ duration: 200 }}
			class="grid flex-1 w-full h-full grid-cols-3 gap-3 p-3 pb-20 place-content-start xl:grid-cols-5 lg:grid-cols-4"
		>
			{#each sticky_notes as sticky_note}
				<Note {categories} {sticky_note} {openEditor} />
			{/each}
			{#if data.view !== 'all' && data.view !== 'recent'}
				<AddNewNote {openModal} />
			{/if}
		</div>
	{/if}
</main>
