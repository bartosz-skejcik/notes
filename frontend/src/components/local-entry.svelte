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
	import type { Notebook } from '../routes/notebooks/[slug]/+page.server';

	type Props = {
		localEntries: LocalEntryType[];
		addChildEntryPending: boolean | null;
		entries: Entry[];
		pendingAIResponseAccept: boolean | null;
		i: number;
		e: LocalEntryType;
		streamResponse: () => void;
		tags: Tag[] | null;
		sessionId: string;
		slug: string | undefined;
		entry: Entry;
		notebook: Notebook;
	};

	let {
		localEntries,
		addChildEntryPending = $bindable(),
		entries = $bindable(),
		i,
		e,
		pendingAIResponseAccept,
		streamResponse,
		tags,
		entry = $bindable(),
		sessionId,
		slug,
		notebook
	}: Props = $props();

	let selectedTag = $derived(tags && tags.filter((t) => t.id === e.tag_id)[0]);
	let sTag = $state<Tag | undefined>();

	let editEntryPending = $state<boolean | null>(false);

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

	function deleteParentEntry(id: number) {
		const entryToDelete = localEntries.find((e) => e.id === id);

		if (entryToDelete) {
			localEntries.filter((e) => e.id !== id);
			console.log('deleted entry', id);
		} else {
			console.log('entries', localEntries);
			console.log('entry not found', id);
		}
	}
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

<div class={`flex w-full pt-2 relative z-[60]`}>
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
				bind:entries
				bind:entry
				childEntry={typeof e.parent_entry_id === 'number'}
				{localEntries}
				onClose={() => (editEntryPending = false)}
				{sessionId}
				{slug}
				{tags}
				{notebook}
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
