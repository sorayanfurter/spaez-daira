
<script lang="ts">
  
  export let input: string;
  export let result: number;

  $: input
    $: result
</script>

<div class="bg-black text-white p-4 text-right">
  <p>{input}</p>
  <p>{result}</p>
</div>

<style>
  p{
    width: 100%;
  }
</style>

