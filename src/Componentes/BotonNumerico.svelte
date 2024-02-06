
<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { input } from './store';

  export let value: number;

  const dispatch = createEventDispatcher();

  function handleClick() {
    input.update(currentValue => `${currentValue}${value}`);
    dispatch('buttonClick', value);
  }
</script>

<button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" on:click={handleClick}>{value}</button>
