<!-- file: src/components/entry.svelte -->

<script lang="ts">
	import { addChildEntry, evalTimePassed, fetchChildEntries } from '$lib/entry';
	import type { Entry } from '$lib/entry';
	import { Check, FlipHorizontal2, Stars, X } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { Button } from '$ui/button';
	import Needle from '$ui/icons/needle.svelte';

	type Props = {
		entries: Entry[];
		entry: Entry;
		slug: string | undefined;
		sessionId: string;
	};

	let pendinAIResponseAccept = $state<boolean | null>(null);

	let { entries = $bindable(), entry, ...data }: Props = $props();

	let localEntries = $state<
		{
			id: number;
			role: string;
			content: string;
			timestamp?: string;
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
				id: Math.floor(Math.random() * 10000),
				role: 'assistant',
				content: ''
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
	}[] {
		const result: {
			id: number;
			role: string;
			content: string;
			timestamp?: string;
		}[] = [];

		function processEntry(entry: Entry) {
			result.push({
				id: entry.id,
				role: entry.role,
				content: entry.title,
				timestamp: entry.timestamp
			});

			// Process children recursively
			if (entry.children && entry.children.length > 0) {
				entry.children.forEach(processEntry);
			}
		}

		processEntry(e);

		return result;
	}

	onMount(async () => {
		localEntries = convertEntriesToLocal(entry);
	});
</script>

<div class="flex flex-col items-start justify-center w-full group">
	{#each localEntries as e, i}
		<div class={`flex w-full pt-2`}>
			<div class="flex flex-col items-center w-10 gap-2">
				<button
					class="flex items-center justify-center rounded-full w-7 h-7 dark:bg-muted bg-muted-foreground/50 aspect-square"
				>
					{#if e.role === 'assistant'}
						<Stars class="w-3/5 h-3/5 text-white/40 text-muted-foreground" />
					{/if}
				</button>
				{#if i !== localEntries.length - 1}
					<div
						class="w-[4px] rounded-full flex-1 min-h-5 dark:bg-muted/80 bg-muted-foreground/20"
					></div>
				{/if}
			</div>
			<div id={e.id.toString()} class="flex-1">
				<p class="pr-2 text-lg">{e.content}</p>
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
	<!-- hover options -->
	<div
		class={`${pendinAIResponseAccept !== null ? 'hidden' : 'flex'}  items-center w-full gap-2 pl-10 mt-2 transition-all duration-200 opacity-0 group-hover:opacity-100 delay-200`}
	>
		<Button variant="brand_ghost" size="sm" class="px-3">
			<Needle class="w-6 h-6 mr-2 text-brand" />
			Add another entry
		</Button>
		{#if !pendinAIResponseAccept}
			<p class="font-extrabold text-center text-muted-foreground/20">/</p>
			<Button variant="brand_ghost" size="sm" class="px-3" onclick={streamResponse}>
				<FlipHorizontal2 class="w-5 h-5 mr-2" />
				Reflect
			</Button>
		{/if}
	</div>
</div>
