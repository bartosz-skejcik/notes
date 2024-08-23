<script lang="ts">
	import * as Popover from '$ui/popover';
	import { Button } from '$ui/button';
	import { colors } from '$lib/utils';

	type Props = {
		selectedColor: string;
		selectColor: (color: string) => void;
	};

	let { selectedColor = $bindable(), selectColor }: Props = $props();
</script>

<Popover.Root>
	<Popover.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="ghost"
			style="background-color: {selectedColor}"
			class="rounded-full w-7 h-7"
			size="icon"
		></Button>
	</Popover.Trigger>
	<Popover.Content class="w-80">
		<div class="grid gap-4">
			<div class="space-y-2">
				<h4 class="font-medium leading-none">Color Picker</h4>
				<p class="text-sm text-muted-foreground">
					Choose one of {colors.length} predefined colors.
				</p>
			</div>
			<div class="flex flex-wrap gap-2">
				{#each colors as color}
					<button
						class="rounded-full w-7 h-7"
						style="background-color: {color}"
						onclick={() => selectColor(color)}
					></button>
				{/each}
			</div>
		</div>
	</Popover.Content>
</Popover.Root>
