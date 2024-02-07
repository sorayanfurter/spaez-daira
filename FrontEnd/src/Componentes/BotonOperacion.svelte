
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

<button  class="bg-red-500 hover:bg-red-700 text-white font-bold py-6 px-6 h-20 w-20  flex items-center justify-center" on:click={handleClick}>{operation}</button>
