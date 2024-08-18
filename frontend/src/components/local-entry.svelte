<script lang="ts">
	import { Stars } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import { FlipHorizontal2 } from 'lucide-svelte';
	import {
		evalTimePassed,
		updateChildEntry,
		updateEntry,
		type Entry,
		type LocalEntryType
	} from '$lib/entry';
	import Needle from '$ui/icons/needle.svelte';
	import Dropdown from './dropdown.svelte';
	import type { Tag } from '$lib/tags';
	import { fly } from 'svelte/transition';

	type Props = {
		localEntries: LocalEntryType[];
		addChildEntryPending: boolean | null;
		pendingAIResponseAccept: boolean | null;
		i: number;
		e: LocalEntryType;
		streamResponse: () => void;
		tags: Tag[] | null;
		sessionId: string;
		slug: string | undefined;
		entry: Entry;
	};

	let {
		localEntries,
		addChildEntryPending = $bindable(),
		i,
		e,
		pendingAIResponseAccept,
		streamResponse,
		tags,
		entry = $bindable(),
		sessionId,
		slug
	}: Props = $props();

	let selectedTag = $derived(tags && tags.filter((t) => t.id === e.tag_id)[0]);
	let sTag = $state<Tag | undefined>();

	$effect(() => {
		if (sTag) {
			// change the tag_id for the entry for the e.id
			entry.tag_id = sTag!.id;
			// TODO: update entry on the server
			(async () => {
				if (!slug) return;
				await updateEntry(sessionId, slug, entry.id, {
					title: entry.title,
					role: entry.role,
					tag_id: selectedTag!.id
				});
			})();
			if (entry.children.length > 0) {
				console.log('entry.children.length > 0');
				entry.children.forEach((ent) => {
					ent.tag_id = sTag!.id;
					// TODO: update child entry on the server
					(async () => {
						if (!slug) return;
						if (!ent.parent_entry_id) return;
						await updateChildEntry(sessionId, slug, ent.parent_entry_id, ent.id, {
							title: ent.title,
							role: ent.role,
							tag_id: selectedTag!.id
						});
					})();
				});
			}
		}
	});
</script>

{#snippet assistantButton(builder: any)}
	<Button
		builders={[builder]}
		size="icon"
		style={selectedTag?.value !== 'none'
			? `background-color: rgb(var(${selectedTag?.color}));`
			: `background-color: hsl(var(${selectedTag?.color}));`}
		class="flex items-center justify-center w-6 h-6 rounded-full aspect-square"
	>
		{#if e.role === 'assistant'}
			<Stars
				class={`w-4/5 h-4/5 text-muted-foreground/40 ${
					selectedTag?.value === 'do-later'
						? `text-green-600`
						: selectedTag?.value === 'highlight'
							? `text-orange-600`
							: selectedTag?.value === 'new-idea'
								? `text-blue-600`
								: `text-muted-foreground/40`
				}`}
			/>
		{/if}
	</Button>
{/snippet}

<div class={`flex w-full pt-2`}>
	<div class="flex flex-col items-center w-10 gap-2 pt-0.5 pr-1">
		<Dropdown button={assistantButton} {tags} bind:e bind:t={sTag} />
		{#if i !== localEntries.length - 1 || addChildEntryPending === true}
			<div
				transition:fly={{ y: -20, duration: 200 }}
				style={selectedTag?.value !== 'none'
					? `background-color: rgb(var(${selectedTag?.color}));`
					: `background-color: hsl(var(${selectedTag?.color}));`}
				class={`w-[4px] ${addChildEntryPending == true ? 'rounded-t-full' : 'rounded-full'} flex-1 min-h-5`}
			></div>
		{/if}
	</div>
	<div id={e.id.toString()} class="flex-1">
		<p class={`pr-2 ${e.role === 'assistant' ? 'text-[0.98rem] text-foreground/80' : 'text-lg'}`}>
			{e.content}
		</p>
		{#if e.role === 'assistant' && localEntries.indexOf(e) !== localEntries.length - 1}
			<br />
		{/if}

		<!-- hover options -->
		<div
			class={`${pendingAIResponseAccept !== null || (localEntries.length > 1 && i !== localEntries.length - 1) ? 'hidden' : 'flex'} ${addChildEntryPending == true ? '' : 'transition-all duration-200 opacity-0 group-hover:opacity-100 delay-200'} items-center w-full gap-2 mt-2`}
		>
			{#if selectedTag}
				<Button
					variant="ghost"
					size="sm"
					onclick={() => (addChildEntryPending = !addChildEntryPending)}
					class={`px-3 ${selectedTag.value == 'do-later' ? 'text-green-600 hover:bg-green-500/30 hover:text-green-600' : ''} ${selectedTag.value == 'highlight' ? 'text-orange-500 hover:bg-orange-500/30 hover:text-orange-500' : ''} ${selectedTag.value == 'new-idea' ? 'hover:bg-brand/30 text-brand hover:text-blue-600' : ''} ${selectedTag.value == 'none' ? 'text-muted-foreground' : ''}`}
				>
					<Needle class="w-6 h-6 mr-2" />
					Add another entry
				</Button>
				{#if !pendingAIResponseAccept}
					<p class="font-extrabold text-center text-muted-foreground/20">/</p>
					<Button
						variant="ghost"
						size="sm"
						class={`px-3 ${selectedTag.value == 'do-later' ? 'text-green-600 hover:bg-green-500/30 hover:text-green-600' : ''} ${selectedTag.value == 'highlight' ? 'text-orange-500 hover:bg-orange-500/30 hover:text-orange-500' : ''} ${selectedTag.value == 'new-idea' ? 'hover:bg-brand/30 text-brand hover:text-blue-600' : ''} ${selectedTag.value == 'none' ? 'text-muted-foreground' : ''}`}
						onclick={() => {
							if (addChildEntryPending === true) {
								addChildEntryPending = false;
							}
							streamResponse();
						}}
					>
						<FlipHorizontal2 class="w-5 h-5 mr-2" />
						Reflect
					</Button>
				{/if}
			{/if}
		</div>
	</div>
	<p class="text-sm text-muted-foreground">
		{e.timestamp && evalTimePassed(e.timestamp)}
	</p>
</div>
