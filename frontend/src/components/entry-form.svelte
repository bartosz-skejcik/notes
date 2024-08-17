<!-- file: src/components/entry-form.svelte -->

<script lang="ts">
	import { Image, X } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import { Label } from '$ui/label';
	import { Textarea } from '$ui/textarea';
	import { autoResize, deleteImage, handleFileUpload } from '$lib/entry-form';
	import type { PageData } from '../routes/notebooks/[slug]/$types';
	import { addEntry, type Entry } from '$lib/entry';

	interface Props extends PageData {
		entries: Entry[];
		parentEntryId?: number;
	}

	let { entries = $bindable(), ...data }: Props = $props();

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
			return;
		}

		const formData = new FormData(e.target as HTMLFormElement);
		const file = formData.get('image_upload') as File;

		const reader = new FileReader();
		reader.onload = async (e) => {
			const img = reader.result as string;

			if (!data.sessionId) return;

			if (!data.slug || typeof data.slug == undefined) return;

			const newEntry = await addEntry(data.sessionId, data.slug, {
				title: title,
				role: 'user'
			});

			if (newEntry !== null) {
				entries.unshift(newEntry);
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
			<div class="p-2 rounded-lg hover:bg-muted text-foreground">
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
		<Button type="submit" variant="default" class="h-8 w-fit">Post</Button>
	</div>
</form>
