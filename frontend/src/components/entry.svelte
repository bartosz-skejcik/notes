<!-- file: src/components/entry.svelte -->

<script lang="ts">
	import { addChildEntry, convertEntriesToLocal } from '$lib/entry';
	import type { Entry } from '$lib/entry';
	import { Check, X } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import EntryForm from './entry-form.svelte';
	import type { Notebook } from '../routes/notebooks/[slug]/+page.server';
	import LocalEntry from './local-entry.svelte';
	import type { Tag } from '$lib/tags';
	import { fly } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import { quintOut } from 'svelte/easing';
	import { useEntriesStore } from '$stores/entries.svelte';

	type Props = {
		entry: Entry;
		sessionId: string;
		slug: string | undefined;
		tags: Tag[] | null;
		notebook: Notebook;
	};

	let { entry, sessionId, slug, tags, notebook }: Props = $props();

	const entriesStore = useEntriesStore();

	let pendingAIResponseAccept = $state<boolean | null>(null);
	let addChildEntryPending = $state<boolean | null>(false);

	let localEntries = $derived(convertEntriesToLocal(entry));

	async function acceptAIResponse() {
		const e = localEntries[localEntries.length - 1];

		if (!slug) {
			console.log(slug);
			return;
		}

		if (!sessionId) {
			console.log(sessionId);
			return;
		}

		await entriesStore.addChildEntry(sessionId, slug, entry.id, {
			title: e.content,
			role: 'assistant',
			tag_id: entry.tag_id,
			parent_entry_id: entry.id
		});

		pendingAIResponseAccept = null;
	}

	function rejectAIResponse() {
		entry.children.pop();
		pendingAIResponseAccept = null;
	}

	async function streamResponse() {
		pendingAIResponseAccept = false;
		let response = '';

		if (!entry.children) {
			console.log('entry.children is null');
			entry.children = [];
		}

		entry.children.push({
			notebook_id: entry.notebook_id,
			id: entry.id + 1,
			role: 'assistant',
			author_id: entry.author_id,
			title: '',
			content: '',
			timestamp: new Date().toISOString(),
			has_photo: false,
			children: [],
			parent_entry_id: entry.id,
			tag_id: entry.tag_id
		});

		const res = await fetch('/api/chat', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				messages: localEntries.map((entry) => {
					return { role: entry.role, content: entry.content };
				})
			})
		});

		const reader = res.body?.getReader();
		const decoder = new TextDecoder();

		while (true) {
			const { value, done } = (await reader?.read()) ?? {};
			if (done) break;
			const chunk = decoder.decode(value, { stream: true });

			// Parse and process the chunk
			const lines = chunk.split('\n');
			for (const line of lines) {
				if (line.startsWith('0:')) {
					// This is a text chunk
					response += JSON.parse(line.slice(2));
					// let lastEntry = entry.children[entry.children.length - 1];
					// lastEntry.title = response;
					// entry.children = [...entry.children];
					entry.children.forEach((e, idx) => {
						if (idx === entry.children.length - 1) {
							e.title = response;
						}
					});
				} else if (line.startsWith('2:')) {
					// This is metadata, you can handle it if needed
					const metadata = JSON.parse(line.slice(2));
					console.log('Metadata:', metadata);
				} else if (line.startsWith('d:')) {
					// This is the end of the stream
					const finalData = JSON.parse(line.slice(2));
					// console.log('Final data:', finalData);
					break;
				}
			}

			await new Promise((resolve) => setTimeout(resolve, 150));
		}
		pendingAIResponseAccept = true;
	}
</script>

<div class="flex flex-col items-start justify-center w-full group">
	{#each localEntries as e, i (i)}
		<div animate:flip={{ duration: 300, easing: quintOut }} class="w-full">
			<LocalEntry
				{localEntries}
				bind:addChildEntryPending
				bind:entry
				{i}
				{e}
				{pendingAIResponseAccept}
				{streamResponse}
				{tags}
				{sessionId}
				{slug}
			/>
		</div>
	{/each}

	<!-- accept / deny AI response -->
	{#if pendingAIResponseAccept === true}
		<div transition:fly={{ y: -20, duration: 200 }}>
			<div class="flex items-center justify-end w-full gap-2 pl-10 mt-3">
				<Button variant="red_ghost" size="sm" class="px-3 h-7" onclick={rejectAIResponse}>
					<X class="w-5 h-5 mr-2" />
					Close
				</Button>
				<Button variant="default" size="sm" class="px-3 h-7" onclick={acceptAIResponse}>
					<Check class="w-5 h-5 mr-2" />
					Save AI response
				</Button>
			</div>
		</div>
	{/if}

	<!-- add another entry -->
	{#if addChildEntryPending === true}
		<div transition:fly={{ y: -20, duration: 300 }} class={`flex items-start w-full gap-2`}>
			<div class="flex flex-col items-center justify-start w-10 gap-2 pb-2 pr-1">
				<div
					style={tags?.filter((t) => t.id === entry.tag_id)[0]?.value !== 'none'
						? `background-color: rgb(var(${tags?.filter((t) => t.id === entry.tag_id)[0]?.color}));`
						: `background-color: hsl(var(${tags?.filter((t) => t.id === entry.tag_id)[0]?.color}));`}
					class="w-[4px] rounded-b-full flex-1 min-h-5"
				></div>
				<button
					style={tags?.filter((t) => t.id === entry.tag_id)[0]?.value !== 'none'
						? `background-color: rgb(var(${tags?.filter((t) => t.id === entry.tag_id)[0]?.color}));`
						: `background-color: hsl(var(${tags?.filter((t) => t.id === entry.tag_id)[0]?.color}));`}
					class="flex items-center justify-center w-6 h-6 rounded-full dark:bg-muted bg-muted-foreground/50 aspect-square"
				>
				</button>
			</div>
			<div class="flex-1 pt-3 -ml-3">
				<EntryForm
					childEntry={true}
					onClose={() => (addChildEntryPending = false)}
					{entry}
					entryId={entry.id}
					{slug}
					{sessionId}
				/>
			</div>
		</div>
	{/if}
</div>
