<script lang="ts">
	import ThemeSwitcher from '$components/theme-switcher.svelte';

	import { Button } from '$ui/button';
	import { ChevronLeft, Home, Settings, Trash } from 'lucide-svelte';
	import { fly } from 'svelte/transition';

	type Props = {
		onClose: (modal: 'new_note' | 'editor') => void;
		view: string;
		open: boolean;
		editorOpen: number | null;
	};

	let { onClose, view, open, editorOpen }: Props = $props();
</script>

<nav
	class={`sticky top-0 flex items-center  px-3 py-2 z-4 bg-background ${open || editorOpen !== null ? 'justify-between' : 'justify-end'}`}
>
	{#if open || editorOpen !== null}
		<div transition:fly={{ x: -20, duration: 200 }}>
			<Button
				variant="ghost"
				size="sm"
				onclick={() => {
					if (open) {
						onClose('new_note');
					}

					if (editorOpen) {
						onClose('editor');
					}
				}}
			>
				<ChevronLeft class="w-4 h-4 mr-2" />
				Back
			</Button>
		</div>
	{/if}
	<div class="flex items-center justify-center gap-4">
		<Button
			variant="red_ghost"
			size="sm"
			class={`border border-red-500/20 ${view == 'recent' || view == 'all' ? 'hidden' : 'flex'}`}
		>
			<Trash class="w-4 h-4 mr-2" />
			Delete category
		</Button>
		<div class="flex items-center justify-center gap-3">
			<ThemeSwitcher />
			<Button size="icon" variant="ghost" class={'relative text-muted-foreground w-8 h-8'}>
				<Settings class="w-5 h-5" />
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
