<!-- file: src/components/entry.svelte -->

<script lang="ts">
	import { addChildEntry, evalTimePassed } from '$lib/entry';
	import type { Entry } from '$lib/entry';
	import { Check, FlipHorizontal2, Stars, X } from 'lucide-svelte';
	import { Button } from '$ui/button';
	import Needle from '$ui/icons/needle.svelte';
	import EntryForm from './entry-form.svelte';
	import type { Notebook } from '../routes/notebooks/[slug]/+page.server';

	type Props = {
		entries: Entry[];
		entry: Entry;
		slug: string | undefined;
		sessionId: string;
		notebook: Notebook;
	};

	let pendinAIResponseAccept = $state<boolean | null>(null);
	let addChildEntryPending = $state<boolean | null>(false);

	let { entries = $bindable(), entry, ...data }: Props = $props();

	let localEntries = $state<
		{
			id: number;
			role: string;
			content: string;
			timestamp?: string;
			parent_entry_id?: number;
		}[]
	>([]);

	async function acceptAIResponse() {
		const e = localEntries[localEntries.length - 1];

		if (!data.slug) {
			console.log(data.slug);
			return;
		}

		if (!data.sessionId) {
			console.log(data.sessionId);
			return;
		}

		await addChildEntry(data.sessionId, data.slug, entry.id, {
			title: e.content,
			role: 'assistant',
			tag_id: entry.tag_id,
			parent_entry_id: entry.id
		});

		pendinAIResponseAccept = null;
	}

	function rejectAIResponse() {
		localEntries.pop();
		localEntries = [...localEntries];
		pendinAIResponseAccept = null;
	}

	async function streamResponse() {
		pendinAIResponseAccept = false;
		let response = '';
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

		localEntries = [
			...localEntries,
			{
				id: entry.id + 1,
				role: 'assistant',
				content: '',
				parent_entry_id: entry.id
			}
		];

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
					let lastEntry = localEntries[localEntries.length - 1];
					lastEntry.content = response;
					localEntries = [...localEntries];
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
		pendinAIResponseAccept = true;

		const lastEntry = localEntries[localEntries.length - 1];
		lastEntry.timestamp = new Date().toISOString();
		localEntries = [...localEntries];
	}

	function convertEntriesToLocal(e: Entry): {
		id: number;
		role: string;
		content: string;
		timestamp?: string;
		parent_entry_id?: number;
	}[] {
		const result: {
			id: number;
			role: string;
			content: string;
			timestamp?: string;
			parent_entry_id?: number;
		}[] = [];

		function processEntry(entry: Entry) {
			result.push({
				id: entry.id,
				role: entry.role,
				content: entry.title,
				timestamp: entry.timestamp,
				parent_entry_id: entry.parent_entry_id
			});

			// Process children recursively
			if (entry.children && entry.children.length > 0) {
				entry.children.forEach(processEntry);
			}
		}

		processEntry(e);

		return result;
	}

	$effect(() => {
		if (entry) {
			localEntries = convertEntriesToLocal(entry);
		}
	});

	$inspect(localEntries);
</script>

<div class="flex flex-col items-start justify-center w-full group">
	{#each localEntries as e, i}
		<div class={`flex w-full pt-2`}>
			<div class="flex flex-col items-center w-10 gap-2 pt-0.5 pr-1">
				<button
					class="flex items-center justify-center w-6 h-6 rounded-full bg-muted-foreground/20 aspect-square"
				>
					{#if e.role === 'assistant'}
						<Stars class="w-4/5 h-4/5 text-muted-foreground/40" />
					{/if}
				</button>
				{#if i !== localEntries.length - 1 || addChildEntryPending === true}
					<div
						class={`w-[4px] ${addChildEntryPending == true ? 'rounded-t-full' : 'rounded-full'} flex-1 min-h-5 dark:bg-muted/80 bg-muted-foreground/20`}
					></div>
				{/if}
			</div>
			<div id={e.id.toString()} class="flex-1">
				<p class="pr-2 text-lg">
					{e.content}
				</p>
				{#if e.role === 'assistant' && localEntries.indexOf(e) !== localEntries.length - 1}
					<br />
				{/if}

				<!-- hover options -->
				<div
					class={`${pendinAIResponseAccept !== null || (localEntries.length > 1 && i !== localEntries.length - 1) ? 'hidden' : 'flex'} ${addChildEntryPending == true ? '' : 'transition-all duration-200 opacity-0 group-hover:opacity-100 delay-200'} items-center w-full gap-2 mt-2`}
				>
					<Button
						variant="brand_ghost"
						size="sm"
						class="px-3"
						onclick={() => (addChildEntryPending = !addChildEntryPending)}
					>
						<Needle class="w-6 h-6 mr-2 text-brand" />
						Add another entry
					</Button>
					{#if !pendinAIResponseAccept}
						<p class="font-extrabold text-center text-muted-foreground/20">/</p>
						<Button
							variant="brand_ghost"
							size="sm"
							class="px-3"
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
				</div>
			</div>
			<p class="text-sm text-muted-foreground">
				{e.timestamp && evalTimePassed(e.timestamp)}
			</p>
		</div>
	{/each}

	<!-- accept / deny AI response -->
	{#if pendinAIResponseAccept === true}
		<div class="flex items-center justify-end w-full gap-2 pl-10 mt-3">
			<Button variant="red_ghost" size="sm" class="px-3 h-7" onclick={rejectAIResponse}>
				<X class="w-5 h-5 mr-2" />
				Close
			</Button>
			<Button variant="brand_solid" size="sm" class="px-3 h-7" onclick={acceptAIResponse}>
				<Check class="w-5 h-5 mr-2" />
				Save AI response
			</Button>
		</div>
	{/if}

	<!-- add another entry -->
	{#if addChildEntryPending === true}
		<div class={`flex items-start w-full gap-2`}>
			<div class="flex flex-col items-center justify-start w-10 gap-2 pb-2 pr-1">
				<div
					class="w-[4px] rounded-b-full flex-1 min-h-5 dark:bg-muted/80 bg-muted-foreground/20"
				></div>
				<button
					class="flex items-center justify-center w-6 h-6 rounded-full dark:bg-muted bg-muted-foreground/50 aspect-square"
				>
				</button>
			</div>
			<div class="flex-1 pt-3 -ml-3">
				<EntryForm
					{...data}
					bind:entries
					childEntry={true}
					parentEntry={entry}
					bind:localEntries
					onClose={() => (addChildEntryPending = false)}
				/>
			</div>
		</div>
	{/if}
</div>
