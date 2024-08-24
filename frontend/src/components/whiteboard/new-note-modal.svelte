<script lang="ts">
	import { Button } from '$ui/button';
	import { fade } from 'svelte/transition';
	import Markdown from 'svelte-exmarkdown';
	import { gfmPlugin } from 'svelte-exmarkdown/gfm';

	const plugins = [gfmPlugin()];

	type Props = {
		onSubmit: (e: Event) => void;
		title: string;
		content: string;
	};

	let { onSubmit, content = $bindable(), title = $bindable() }: Props = $props();
</script>

<form
	onsubmit={onSubmit}
	transition:fade={{ duration: 200 }}
	class="z-[80] w-full flex flex-col items-start justify-start gap-8 p-8"
>
	<input
		type="text"
		name="title"
		id="title"
		bind:value={title}
		placeholder="Write a note title"
		autocomplete="off"
		class="w-full text-2xl bg-transparent focus-visible:outline-none placeholder:text-muted-foreground"
	/>
	<div class="flex flex-1 w-full h-full">
		<textarea
			name="content"
			id="content"
			bind:value={content}
			placeholder="Here goes your note content"
			autocomplete="off"
			class="flex-1 w-1/2 text-lg bg-transparent focus-visible:outline-none placeholder:text-muted-foreground"
			rows={10}
		></textarea>
		<article
			class="grid w-1/2 h-full grid-cols-1 grid-rows-1 prose place-items-start text-muted-foreground"
		>
			<Markdown md={content} {plugins} />
		</article>
	</div>
	<Button type="submit">Create</Button>
</form>
