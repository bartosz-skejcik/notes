<script lang="ts">
	import { Plus } from 'lucide-svelte';

	type Props = {
		openModal: () => void;
	};

	let props: Props = $props();
</script>

<button
	onclick={props.openModal}
	class="flex flex-col group items-center justify-center w-full h-56 col-span-1 bg-muted-foreground/5 hover:bg-muted-foreground/10 transition-all duration-200 border-2 rounded-lg p-2 border-dotted border-muted-foreground/20"
>
	<Plus class="w-7 h-7 text-muted-foreground group-hover:scale-110 transition-all duration-200" />
</button>
