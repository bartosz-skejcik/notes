<script lang="ts">
	import { Button } from '$components/ui/button';
	import ColorPicker from '$components/whiteboard/color-picker.svelte';
	import { createCategory, type Category } from '$lib/whiteboard/categories';
	import { ChevronRight, X } from 'lucide-svelte';
	import { onMount } from 'svelte';

	type Props = {
		categories: Category[];
		notebookId: string | null;
		sessionId: string | null;
	};

	const { categories = $bindable(), notebookId, sessionId }: Props = $props();

	let name = $state('New category');
	let color = $state('#2563eb');
	let id = categories[categories.length - 1].id + 1;

	async function handleSubmit() {
		if (notebookId && sessionId) {
			const newId = await createCategory(notebookId, name, color, sessionId);
			if (newId) {
				categories.forEach((category) => {
					if (category.id === id) {
						category.name = name;
					}
				});
			}
		}
	}

	onMount(() => {
		categories.push({
			notebookId: Number(notebookId),
			id,
			name,
			color
		});

		return () => {
			if (name == 'New category') {
				categories.pop();
			}
		};
	});

	function selectColor(c: string) {
		color = c;
		categories.forEach((category) => {
			if (category.id === id) {
				category.color = c;
			}
		});
	}
</script>

<div class="flex flex-col items-center justify-center flex-1 w-full h-full">
	<div class="flex items-center justify-center px-2 border-2 rounded-xl border-muted">
		<ColorPicker bind:selectedColor={color} {selectColor} />
		<input
			type="text"
			bind:value={name}
			placeholder="New category"
			class="flex-1 w-full pb-2 pt-1.5 ml-3 bg-transparent text-start focus:outline-none text-foreground/80"
		/>
	</div>
	<p class="w-64 mt-4 text-center text-muted-foreground">
		This category is empty. Start by adding a new sticky note.
	</p>
	<div class="w-1 h-24 mt-4 rounded-full bg-foreground/20"></div>
	<div class="flex items-center justify-center gap-4 mt-5">
		<Button variant="red_ghost" size="sm" onclick={() => categories.pop()}>
			<X class="w-4 h-4" />
			Delete
		</Button>
		<Button variant="default" size="sm" onclick={handleSubmit}>
			Submit
			<ChevronRight class="w-4 h-4 ml-2" />
		</Button>
	</div>
</div>
