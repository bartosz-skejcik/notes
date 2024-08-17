<!-- file: src/components/entry-form.svelte -->

<script lang="ts">
	import { Check, Image, X } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import { Label } from '$ui/label';
	import { Textarea } from '$ui/textarea';
	import { autoResize, deleteImage, handleFileUpload } from '$lib/entry-form';
	import type { PageData } from '../routes/notebooks/[slug]/$types';
	import { addChildEntry, addEntry, type Entry } from '$lib/entry';

	interface Props extends PageData {
		childEntry?: boolean;
		onClose?: () => void;
		entries: Entry[];
		parentEntry?: Entry;
		localEntries?: {
			id: number;
			role: string;
			content: string;
			timestamp?: string;
			parent_entry_id?: number;
		}[];
	}

	let {
		parentEntry,
		childEntry,
		onClose,
		entries = $bindable(),
		localEntries = $bindable(),
		...data
	}: Props = $props();

	let title = $state('');
	let file = $state<File | undefined>(undefined);

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

			if (childEntry && parentEntry && parentEntry.id) {
				const newEntry = await addChildEntry(data.sessionId, data.slug, parentEntry.id, {
					title: title,
					role: 'user',
					tag_id: parentEntry.tag_id ?? null
				});

				console.log(newEntry);

				// {
				// 	id: 20,
				// 	notebook_id: 1,
				// 	author_id: 1,
				// 	title: 'this is a test string to see if the child entry adding works',
				// 	content: '',
				// 	role: 'user',
				// 	timestamp: '2024-08-17T11:42:17.254088Z',
				// 	parent_entry_id: 18,
				// 	tag_id: null
				// }

				if (newEntry !== null) {
					// entries.unshift(newEntry);
					localEntries?.push({
						id: newEntry.id,
						role: newEntry.role,
						content: newEntry.title,
						timestamp: newEntry.timestamp,
						parent_entry_id: newEntry.parent_entry_id
					});
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
			setImagePreview(null);
			setImagePreviewDims(null);
		};

		reader.readAsDataURL(file);
	}
</script>

<form
	onsubmit={(e: Event) => {
		e.preventDefault();
		handleSubmit(e);
	}}
	class="w-full py-2 mb-2"
>
	<div class="w-full">
		<Textarea
			bind:value={title}
			name="title"
			id="title"
			style="word-wrap: break-word; height: {textAreaHeight};"
			class="w-3/4 px-0 ml-2 overflow-hidden text-lg border-none shadow-none resize-none focus-visible:ring-0 whitespace-break-spaces"
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
	<div class="flex items-center justify-between w-full 2xl:justify-start 2xl:gap-4">
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
			{#if childEntry}
				<Button variant="red_ghost" size="sm" class="px-3 h-7" onclick={onClose}>
					<X class="w-5 h-5 mr-2" />
					Close
				</Button>
			{/if}
			<Button
				type="submit"
				variant="default"
				size={childEntry ? 'sm' : 'default'}
				class={`${childEntry ? 'h-7 px-3' : 'h-8'} w-fit`}
			>
				{#if childEntry}
					<Check class="w-5 h-5 mr-2" />
					Reply
				{:else}
					Post
				{/if}
			</Button>
		</div>
	</div>
</form>
