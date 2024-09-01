<!-- file: src/components/local-entry.svelte -->

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
	import EntryForm from './entry-form.svelte';
	import { useEntriesStore } from '$stores/entries.svelte';

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

	let entriesStore = useEntriesStore();

	let selectedTag = $derived(tags && tags.filter((t) => t.id === e.tag_id)[0]);
	let sTag = $state<Tag | undefined>();

	let editEntryPending = $state<boolean | null>(false);

	$effect(() => {
		if (sTag) {
			(async () => {
				await updateEntryTag(entry, sTag);
				console.log('updated entry tag');
			})();
		}
	});

	async function updateEntryTag(entryToUpdate: Entry, newTag: Tag) {
		if (!slug) return;

		entriesStore.entries.forEach(async (entry) => {
			if (entry.id === entryToUpdate.id) {
				// Update parent entry
				entry.tag_id = newTag.id;
				await updateEntry(sessionId, slug, entry.id, {
					title: entry.title,
					role: entry.role,
					tag_id: newTag.id
				});

				// Update child entries
				if (entry.children && entry.children.length > 0) {
					for (const childEntry of entry.children) {
						childEntry.tag_id = newTag.id;
						await updateChildEntry(sessionId, slug, childEntry.parent_entry_id!, childEntry.id, {
							title: childEntry.title,
							role: childEntry.role,
							tag_id: newTag.id
						});
					}
				}
			}
		});
	}

	function deleteParentEntry(id: number) {
		editEntryPending = false;
		if (slug) {
			entriesStore.deleteEntry(sessionId, slug, id);
		}
	}

	function getTagColor(tag: Tag | undefined) {
		return tag?.value !== 'none'
			? `background-color: rgb(var(${tag?.color}));`
			: `background-color: hsl(var(${tag?.color}));`;
	}

	function getTagTextColor(tagValue: string | undefined) {
		switch (tagValue) {
			case 'do-later':
				return 'text-green-600';
			case 'highlight':
				return 'text-orange-600';
			case 'new-idea':
				return 'text-blue-600';
			default:
				return 'text-muted-foreground/40';
		}
	}
</script>

{#snippet assistantButton(builder: any)}
	<Button
		builders={[builder]}
		size="icon"
		style={getTagColor(selectedTag!)}
		class="flex items-center justify-center w-6 h-6 rounded-full aspect-square"
	>
		{#if e.role === 'assistant'}
			<Stars
				class={`w-4/5 h-4/5 text-muted-foreground/40 ${getTagTextColor(selectedTag?.value)}`}
			/>
		{/if}
	</Button>
{/snippet}

<div class={`flex w-full pt-2 relative z-[60]`}>
	<div class="flex flex-col items-center w-10 gap-2 pt-0.5 pr-1">
		<Dropdown button={assistantButton} {tags} bind:e bind:t={sTag} />
		{#if i !== localEntries.length - 1 || addChildEntryPending === true}
			<div
				transition:fly={{ y: -20, duration: 200 }}
				style={getTagColor(selectedTag!)}
				class={`w-[4px] ${addChildEntryPending == true ? 'rounded-t-full' : 'rounded-full'} flex-1 min-h-5`}
			></div>
		{/if}
	</div>
	<div id={e.id.toString()} class="flex-1">
		{#if editEntryPending === false}
			<p
				class={`pr-32 ${e.role === 'assistant' ? 'text-[0.98rem] text-foreground/80 mt-0.5' : 'text-lg'}`}
			>
				{#if e.role === 'assistant'}
					{#if e.content.length > 0}
						{e.content}
					{:else}
						<span class="text-muted-foreground/40 animate-pulse">AI is thinking...</span>
					{/if}
				{:else}
					{e.content}
				{/if}
			</p>
			{#if e.role === 'assistant' && localEntries.indexOf(e) !== localEntries.length - 1}
				<br />
			{/if}
		{:else}
			<EntryForm
				{entry}
				childEntry={typeof e.parent_entry_id === 'number'}
				onClose={() => (editEntryPending = false)}
				{sessionId}
				{slug}
				{editEntryPending}
				content={e.content}
				entryId={e.id}
				onParentEntryDelete={deleteParentEntry}
			/>
		{/if}

		<!-- hover options -->
		<div
			class={`${pendingAIResponseAccept !== null || editEntryPending === true || (localEntries.length > 1 && i !== localEntries.length - 1) ? 'hidden' : 'flex'} ${addChildEntryPending == true ? '' : 'transition-all duration-200 opacity-0 group-hover:opacity-100 delay-200'} items-center w-full gap-2 mt-2`}
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
	<button
		class="absolute top-0 right-0 px-2 py-px text-sm font-medium transition-all duration-200 ease-in-out rounded-lg text-muted-foreground hover:bg-muted-foreground/10 hover:text-foreground h-fit"
		onclick={() => (editEntryPending = !editEntryPending)}
	>
		{e.timestamp && evalTimePassed(e.timestamp)}
	</button>
</div>
