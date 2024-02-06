
<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { input } from './store';

  export let operation: string;

  const dispatch = createEventDispatcher();

  function handleClick() {
    input.update(currentValue => `${currentValue}${operation}`);
    dispatch('buttonClick', operation);
  }
</script>

<button on:click={handleClick}>{operation}</button>
