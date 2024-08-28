<script lang="ts">
	import type { Category } from '$lib/whiteboard/categories';
	import * as Card from '$ui/card';
	import { hexToRGBA } from '$lib/whiteboard/utils';
	import type { StickyNote } from '$lib/whiteboard/sticky_notes';
	import Markdown from 'svelte-exmarkdown';

	type Props = {
		categories: Category[];
		sticky_note: StickyNote;
		openEditor: (note_id: number) => void;
	};

	let { categories, sticky_note, openEditor }: Props = $props();

	let category = $derived(
		categories.find((c) => c.id.toString() === sticky_note.category_id.toString())!
	);
</script>

<button onclick={() => openEditor(sticky_note.id)}>
	<Card.Root
		style={`background: ${hexToRGBA(category.color, 0.7)}; border-color: ${category.color};`}
		class="flex flex-col items-start justify-between w-full h-56 col-span-1"
	>
		<Card.Header class="px-4 py-3 flex items-start text-start">
			<div
				class="flex items-center justify-start gap-1.5 mb-1 text-[0.8rem] font-medium text-center text-foreground/80"
			>
				<p class="px-1.5 py-0.5 font-medium rounded-lg bg-foreground/10">
					{new Date(sticky_note.created_at).toLocaleString('en-PL', {
						year: 'numeric'
					})}
				</p>
				<span class="mb-0.5">•</span>
				<p>{category.name}</p>
			</div>
			<Card.Title class="text-lg">{sticky_note.title}</Card.Title>
			<Card.Description class="prose text-foreground">
				<Markdown
					md={sticky_note.content.length > 50
						? sticky_note.content.slice(0, 50) + '...'
						: sticky_note.content}
				/>
			</Card.Description>
		</Card.Header>
		<Card.Footer class="px-4 py-3">
			<p class="text-xs text-foreground/70">
				{new Date(sticky_note.created_at).toLocaleString('en-PL', {
					hour: 'numeric',
					minute: 'numeric',
					weekday: 'short',
					day: 'numeric',
					month: 'short'
				})}
			</p>
		</Card.Footer>
	</Card.Root>
</button>
