<script lang="ts">
    import {toggleMode} from "mode-watcher";
    import { Button } from "$ui/button";

    import {Sun, Moon} from "lucide-svelte/icons"

    export let className: string = "";
</script>

<Button on:click={toggleMode} class={className + " ml-10 aspect-square w-10 relative"}>
    <Sun
    class="absolute h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
  />
  <Moon
    class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
  />
  <span class="sr-only">Toggle theme</span>
</Button>