<!-- file: src/components/entry-form.svelte -->

<script lang="ts">
	import { Check, Image, X } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import { Label } from '$ui/label';
	import { Textarea } from '$ui/textarea';
	import { autoResize, deleteImage, handleFileUpload } from '$lib/entry-form';
	import type { PageData } from '../routes/notebooks/[slug]/$types';
	import {
		addChildEntry,
		addEntry,
		deleteEntry,
		updateChildEntry,
		updateEntry,
		type Entry,
		type LocalEntryType
	} from '$lib/entry';
	import { fly } from 'svelte/transition';
	import Page from '../routes/+page.svelte';

	interface Props extends PageData {
		childEntry?: boolean;
		onClose?: () => void;
		entries: Entry[];
		entry?: Entry;
		localEntries?: LocalEntryType[];
		editEntryPending?: boolean | null;
		content?: string;
		entryId?: number;
		onParentEntryDelete?: (id: number) => void;
	}

	let {
		childEntry,
		onClose,
		entries = $bindable(),
		localEntries,
		entry = $bindable(),
		editEntryPending,
		content,
		entryId,
		onParentEntryDelete,
		...data
	}: Props = $props();

	let title = $state('');
	let file = $state<File | undefined>(undefined);

	$effect(() => {
		if (content !== undefined) {
			title = content;
		}
	});

	let textAreaHeight = $state('auto');
	let imagePreview = $state<string | null>(null);
	let imagePreviewDims = $state<{
		width: number;
		height: number;
	} | null>(null);

	$effect(() => {
		// Trigger resize whenever title changes
		if (title !== undefined) {
			autoResize();
		}
	});

	function setFile(fileInput: File | undefined) {
		file = fileInput;
	}

	function setImagePreview(imagePrev: string | null) {
		imagePreview = imagePrev;
	}

	function setImagePreviewDims(imagePrevDims: { width: number; height: number } | null) {
		imagePreviewDims = imagePrevDims;
	}

	async function handleSubmit(e: Event) {
		if (!title) {
			alert('Please enter a title');
			return;
		}

		const formData = new FormData(e.target as HTMLFormElement);
		const file = formData.get('image_upload') as File;

		const reader = new FileReader();
		reader.onload = async (e) => {
			const img = reader.result as string;

			if (!data.sessionId) return;

			if (!data.slug || typeof data.slug == undefined) return;

			if (!entry) return;

			if (entry.id) {
				if (editEntryPending == true && title !== undefined && entryId) {
					// // await updateEntry(data.sessionId, data.slug, entry.id, {
					// // 	title: title,
					// // 	role: 'user',
					// // 	tag_id: entry.tag_id ?? null
					// // });

					if (childEntry) {
						console.log('updating child entry', childEntry, entryId);

						const childEntryObject = entry.children.find((ent) => ent.id === entryId);

						await updateChildEntry(data.sessionId, data.slug, entry.id, entryId, {
							...childEntryObject,
							title: title
						});

						entry.children.forEach((ent) => {
							if (ent.id === entryId) {
								ent.title = title;
							}
						});
					} else {
						console.log('updating entry', entry, entryId);
						await updateEntry(data.sessionId, data.slug, entry.id, {
							title: title
						});

						entry.title = title;
					}
				} else if (childEntry) {
					console.log('adding child entry', childEntry);
					const newEntry = await addChildEntry(data.sessionId, data.slug, entry.id, {
						title: title,
						role: 'user',
						tag_id: entry.tag_id ?? null
					});

					if (newEntry !== null) {
						if (!entry.children) {
							entry.children = [];
						}

						entry.children.push({
							notebook_id: entry.notebook_id,
							id: newEntry.id,
							role: newEntry.role,
							author_id: entry.author_id,
							title: newEntry.title,
							content: newEntry.content,
							timestamp: newEntry.timestamp,
							has_photo: false,
							children: [],
							parent_entry_id: newEntry.parent_entry_id,
							tag_id: newEntry.tag_id
						});
					}
				}
			} else {
				console.log('adding entry');
				const newEntry = await addEntry(data.sessionId, data.slug, {
					title: title,
					role: 'user'
				});

				if (newEntry !== null) {
					entries.unshift(newEntry);
				}
			}

			title = '';
			setFile(undefined);
			onClose && onClose();
			setImagePreview(null);
			setImagePreviewDims(null);
		};

		reader.readAsDataURL(file);
	}

	function animate(node: HTMLElement, args: any): any {
		if (editEntryPending === true) {
			return fly(node, args);
		} else {
			return fly(node, args);
		}
	}
</script>

<form
	onsubmit={(e: Event) => {
		e.preventDefault();
		handleSubmit(e);
	}}
	class={`w-full mb-2 ${editEntryPending ? 'pb-2' : 'py-2'}`}
>
	<div class="w-full">
		<Textarea
			bind:value={title}
			name="title"
			id="title"
			style="word-wrap: break-word; height: {textAreaHeight};"
			class={`w-3/4 px-0 overflow-hidden text-lg border-none shadow-none resize-none focus-visible:ring-0 whitespace-break-spaces ${editEntryPending ? 'ml-0 py-0' : 'ml-2'}`}
			placeholder="What are you thinking?"
			rows={1}
		/>
	</div>
	{#if imagePreview}
		<div class="relative mt-2 w-44">
			<img
				src={imagePreview}
				alt="Preview"
				class="object-cover rounded-lg"
				style={`width: ${imagePreviewDims?.width}px; height: ${imagePreviewDims?.height}px;`}
			/>
			<button
				type="button"
				class="absolute p-1 text-white bg-black bg-opacity-50 rounded-full top-1 right-1"
				onclick={() => deleteImage(setFile, setImagePreview)}
			>
				<X class="w-4 h-4" />
			</button>
		</div>
	{/if}
	<div
		transition:animate={{ y: -20, duration: 300 }}
		class="flex items-center justify-between w-full 2xl:justify-start 2xl:gap-4"
	>
		<Label for="image_upload" class="cursor-pointer">
			<div class="p-1 rounded-lg hover:bg-muted-foreground/20 text-foreground bg-muted">
				<Image class="w-5 h-5" />
			</div>
		</Label>
		<input
			id="image_upload"
			type="file"
			name="image_upload"
			class="hidden"
			accept="image/*"
			onchange={(event: Event) =>
				handleFileUpload(
					(event.target as HTMLInputElement).files,
					setFile,
					setImagePreview,
					imagePreview,
					setImagePreviewDims
				)}
		/>
		<div class="flex items-center justify-end w-full gap-2">
			{#if childEntry || editEntryPending}
				<Button
					variant="red_ghost"
					size="sm"
					class="px-3 h-7"
					onclick={async () => {
						if (editEntryPending) {
							// delete entry
							if (data.slug && entryId) {
								const deletedEntry = await deleteEntry(data.sessionId, data.slug, entryId);

								if (deletedEntry && deletedEntry.id) {
									// get the index of the deleted entry
									let index = -1;

									if (childEntry && entry) {
										index = entry.children.findIndex((e) => e.id === deletedEntry.id);

										console.log('child, index', index);

										// splice the deleted entry from the entries array
										entry.children.splice(index, 1);

										console.log('deleted entry', deletedEntry.id, index);
									} else {
										onParentEntryDelete && onParentEntryDelete(deletedEntry.id);
									}
								}
							}
						} else {
							onClose && onClose();
						}
					}}
				>
					<X class="w-5 h-5 mr-2" />
					{#if editEntryPending}
						Delete
					{:else}
						Close
					{/if}
				</Button>
			{/if}
			<Button
				type="submit"
				variant="default"
				size={childEntry || editEntryPending ? 'sm' : 'default'}
				class={`${childEntry || editEntryPending ? 'h-7 px-3' : 'h-8'} w-fit`}
			>
				{#if editEntryPending}
					<Check class="w-5 h-5 mr-2" />
					Update
				{:else if childEntry}
					<Check class="w-5 h-5 mr-2" />
					Reply
				{:else}
					Post
				{/if}
			</Button>
		</div>
	</div>
</form>
