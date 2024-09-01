<!-- file: src/components/entry-form.svelte -->

<script lang="ts">
	import { Check, Image, X } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import { Label } from '$ui/label';
	import { Textarea } from '$ui/textarea';
	import { autoResize, deleteImage, handleFileUpload } from '$lib/entry-form';
	import type { PageData } from '../routes/notebooks/[slug]/$types';
	import type { Entry } from '$lib/entry';
	import { fly } from 'svelte/transition';
	import { useEntriesStore } from '$stores/entries.svelte';

	interface Props extends PageData {
		childEntry?: boolean;
		onClose?: () => void;
		entry?: Entry;
		editEntryPending?: boolean | null;
		content?: string;
		entryId?: number;
		onParentEntryDelete?: (id: number) => void;
	}

	let {
		childEntry,
		onClose,
		entry,
		editEntryPending,
		content,
		entryId,
		onParentEntryDelete,
		...data
	}: Props = $props();

	const entriesStore = useEntriesStore();

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
		e.preventDefault();
		if (!title) {
			alert('Please enter a title');
			return;
		}

		const formData = new FormData(e.target as HTMLFormElement);
		const file = formData.get('image_upload') as File;

		if (!data.sessionId || !data.slug) {
			console.log('Missing session ID or slug');
			return;
		}

		console.log(
			`editEntryPending ${editEntryPending}`,
			`childEntry ${childEntry}`,
			`entry ${entry}`,
			`entryId ${entryId}`
		);

		if (editEntryPending && entryId) {
			if (childEntry && entry) {
				console.log('updateChildEntry');
				await entriesStore.updateChildEntry(data.sessionId, data.slug, entry.id, entryId, {
					title
				});
			} else {
				console.log('updateEntry');
				await entriesStore.updateEntry(data.sessionId, data.slug, entryId, { title });
			}
		} else if (childEntry && entry) {
			console.log('addChildEntry');
			await entriesStore.addChildEntry(data.sessionId, data.slug, entry.id, {
				title,
				role: 'user',
				tag_id: entry.tag_id
			});
		} else {
			console.log('addEntry');
			await entriesStore.addEntry(data.sessionId, data.slug, { title, role: 'user' });
		}

		title = '';
		setFile(undefined);
		onClose && onClose();
		setImagePreview(null);
		setImagePreviewDims(null);
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
								await entriesStore.deleteEntry(data.sessionId, data.slug, entryId);
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
