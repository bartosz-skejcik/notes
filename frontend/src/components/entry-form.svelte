<script lang="ts">
	import { Image, X } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import { Label } from '$ui/label';
	import { Textarea } from '$ui/textarea';
	import { Input } from '$ui/input';
	import { autoResize, deleteImage, handleFileUpload } from '$lib/entry-form';

	let content = $state('');
	let file = $state<File | null>(null);

	let textAreaHeight = $state('auto');
	let imagePreview = $state<string | null>(null);
	let imagePreviewDims = $state<{
		width: number;
		height: number;
	} | null>(null);

	$effect(() => {
		// Trigger resize whenever content changes
		if (content !== undefined) {
			autoResize();
		}
	});

	function setFile(fileInput: File | null) {
		file = fileInput;
	}

	function setImagePreview(imagePrev: string | null) {
		imagePreview = imagePrev;
	}

	function setImagePreviewDims(imagePrevDims: { width: number; height: number } | null) {
		imagePreviewDims = imagePrevDims;
	}
</script>

<form method="POST" action="/api/entries" class="w-full py-2 mb-2">
	<div class="w-full">
		<Textarea
			bind:value={content}
			name="content"
			style="word-wrap: break-word; height: {textAreaHeight};"
			class="w-3/4 px-0 overflow-hidden text-lg border-none shadow-none resize-none focus-visible:ring-0 whitespace-break-spaces"
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
		<Label for="image-upload" class="cursor-pointer">
			<div class="p-2 rounded-lg hover:bg-muted text-foreground">
				<Image class="w-5 h-5" />
			</div>
		</Label>
		<Input
			id="image-upload"
			type="file"
			name="image"
			class="hidden"
			accept="image/*"
			onchange={(event: Event) =>
				handleFileUpload(event, setFile, setImagePreview, imagePreview, setImagePreviewDims)}
		/>
		<Button type="submit" variant="default" class="h-8 w-fit">Post</Button>
	</div>
</form>
