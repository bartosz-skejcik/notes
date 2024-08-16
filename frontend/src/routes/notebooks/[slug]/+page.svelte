<script lang="ts">
	import { createSearch } from '$stores/search.svelte';

	import CommandSearch from '$components/command-search.svelte';
	import ThemeSwitcher from '$components/theme-switcher.svelte';
	import { Button } from '$ui/button';

	import { Settings, Home } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import EntryForm from '$components/entry-form.svelte';

	let { data } = $props();

	const search = createSearch();

	const items = [
		{ label: 'New Notebook', value: 'new' },
		{ label: 'Open Notebook', value: 'open' },
		{ label: 'Save Notebook', value: 'save' },
		{ label: 'Save Notebook As', value: 'save-as' },
		{ label: 'Rename Notebook', value: 'rename' },
		{ label: 'Delete Notebook', value: 'delete' },
		{ label: 'Export Notebook', value: 'export' },
		{ label: 'Import Notebook', value: 'import' },
		{ label: 'Print Notebook', value: 'print' },
		{ label: 'Print Notebook Preview', value: 'print-preview' },
		{ label: 'Exit Notebook', value: 'exit' }
	];

	let messages = $state([
		{
			role: 'user',
			content: `The CEO of Figma reposted my design.\nMy life is complete 😌`
		}
	]);

	let response = $state('');

	async function streamResponse() {
		response = '';
		messages = [...messages, { role: 'assistant', content: '' }];
		const res = await fetch('/api/chat', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ messages })
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
					let lastMessage = messages[messages.length - 1];
					lastMessage.content = response;
					messages = messages;
				} else if (line.startsWith('2:')) {
					// This is metadata, you can handle it if needed
					const metadata = JSON.parse(line.slice(2));
					console.log('Metadata:', metadata);
				} else if (line.startsWith('d:')) {
					// This is the end of the stream
					const finalData = JSON.parse(line.slice(2));
					console.log('Final data:', finalData);
					break;
				}
			}
			await new Promise((resolve) => setTimeout(resolve, 50));
		}
	}

	// onMount(async () => {
	// 	await streamResponse();
	// });
</script>

<nav class="flex items-center justify-between px-3 py-2">
	<div class="flex items-center gap-2 text-[0.97rem] text-muted-foreground/80">
		<p>{data.notebook.name}</p>
		<span>•</span>
		<p>
			{new Date().toLocaleString('en-PL', {
				weekday: 'long',
				year: 'numeric',
				month: 'long',
				day: 'numeric'
			})}
		</p>
	</div>
	<div class="flex items-center justify-end w-1/2 gap-4">
		<CommandSearch {items} onChange={search.setValue} />
		<div class="flex items-center justify-center gap-3">
			<ThemeSwitcher />
			<Button size="icon" variant="ghost" class="w-8 h-8">
				<Settings class="w-5 h-5 text-muted-foreground" />
			</Button>
			<Button
				onclick={() => (window.location.href = '/')}
				size="icon"
				variant="ghost"
				class="w-8 h-8"
			>
				<Home class="w-5 h-5 text-muted-foreground" />
			</Button>
		</div>
	</div>
</nav>
<div class="p-3">
	<EntryForm />
	{#each messages as message}
		<p class="text-lg">{message.content}</p>
	{/each}
</div>
