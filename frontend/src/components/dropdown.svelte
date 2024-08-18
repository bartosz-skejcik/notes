<script lang="ts">
	import * as Popover from '$ui/popover';
	import type { Tag } from '$lib/tags';
	import type { LocalEntryType } from '$lib/entry';

	type Props = {
		button: any;
		tags: Tag[] | null;
		e: LocalEntryType | undefined;
		t: Tag | undefined | null;
	};

	let { button, e = $bindable(), t = $bindable(), tags }: Props = $props();
</script>

<Popover.Root>
	<Popover.Trigger asChild let:builder>
		{@render button(builder)}
	</Popover.Trigger>
	<Popover.Content class="w-40 p-1">
		{#if tags}
			{#each tags as tag}
				<button
					onclick={() => {
						if (e) {
							e.tag_id = tag.id;
						}
						t = tag;
					}}
					class="flex items-center justify-start w-full gap-2 px-2 py-0.5 text-lg transition-all duration-200 border border-transparent rounded-lg hover:border-border hover:bg-muted/50"
				>
					<div
						class="w-5 h-5 rounded-full"
						style={tag.value !== 'none'
							? `background: rgb(var(${tag.color}));`
							: `background: hsl(var(${tag.color}));`}
					></div>
					{tag.label}
				</button>
			{/each}
		{/if}
	</Popover.Content>
</Popover.Root>
