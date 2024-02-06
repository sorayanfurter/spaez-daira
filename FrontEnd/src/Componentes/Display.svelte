
<script lang="ts">
  
  export let input: string;
  export let result: number;

  $: input
    $: result
</script>

<div>
  <p>Entrada: {input}</p>
  <p>Resultado: {result}</p>
</div>

<style>
  p{
    width: 100%;
  }
</style>

