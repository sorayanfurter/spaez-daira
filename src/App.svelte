<script context="module" lang="ts">
  import type { AxiosResponse } from 'axios';
  import { onMount } from "svelte";
  import type { HistoryEntry } from './Componentes/Historial.svelte';
  import Display from "./Componentes/Display.svelte";
  import BotonNumerico from "./Componentes/BotonNumerico.svelte";
  import BotonOperacion from "./Componentes/BotonOperacion.svelte";
  import Historial from "./Componentes/Historial.svelte";
  import axios from 'axios';

</script>

<script lang="ts">
  let input: string = '';
  let result: number = 0;
  let history: HistoryEntry[] = [];
  $: {
    input;
    result;
    history
  }
 

  function handleNumericInput(value: number): void {
    input += value.toString();
  }

  function handleOperationInput(operation: string): void {
    if (input.length > 0) {
      input += ` ${operation} `;
    }
  }

  async function fetchHistory() {
    try {
      console.log("Fetching history..."); 
      const response: AxiosResponse<HistoryEntry[]> = await axios.get(
        "http://localhost:9090/Historial"
      );
      
    console.log("Response data:", response.data)
      history = response.data;
      
    } catch (error) {
      console.error("Error fetching history:", error);
    }
  }

  onMount(() => {
    fetchHistory();
  });

 

  async function handleButtonClick(event: { detail: string | number }): Promise<void> {
    const clickedValue = event.detail;

    if (clickedValue === "=") {
      if (input.length > 0) {
        try {
          console.log("Fetching history...");
          const response: AxiosResponse<{ resultado: number }> = await axios.post(
            "http://localhost:9090/Calcular",
            {
              n1: parseFloat(input.split(" ")[0]),
              n2: parseFloat(input.split(" ")[2]),
              operacion: input.split(" ")[1],
            }
          );
          console.log("Response status:", response.status);
          console.log("Response data:", response.data);
          result = response.data.resultado;
          
          
          // Reiniciar el input después de mostrar el resultado
          input = "";

          // Actualizar historial
          fetchHistory();
        } catch (error) {
          console.error("Error calculating:", error);
        }
        
      }
    } else if (clickedValue === "C") {
      // Resetear input and result
      input = "";
      result = 0;
    } else if (typeof clickedValue === "number") {
      handleNumericInput(clickedValue);
    } else if (typeof clickedValue === "string") {
      handleOperationInput(clickedValue);
    }

  }

  
</script>

<style>
  @import "tailwindcss/base.css";
  @import "tailwindcss/components.css";
  @import "tailwindcss/utilities.css";
  
</style>
<!-- Display -->
<div>
  <Display bind:input={input} bind:result={result}/>
</div>

<!-- Teclado numérico -->
<div>
  {#each [1, 2, 3, 4, 5, 6, 7, 8, 9, 0] as num}
    <BotonNumerico value={num} on:buttonClick={handleButtonClick} />
  {/each}
</div>

<!-- Teclado operación-->
<div>
  <BotonOperacion operation="+" on:buttonClick={handleButtonClick} />
  <BotonOperacion operation="-" on:buttonClick={handleButtonClick} />
  <BotonOperacion operation="*" on:buttonClick={handleButtonClick} />
  <BotonOperacion operation="/" on:buttonClick={handleButtonClick} />
  <BotonOperacion operation="C" on:buttonClick={handleButtonClick} />
  <BotonOperacion operation= "=" on:buttonClick={handleButtonClick}/>
</div>

<!-- Historial -->
{#if history && history.length > 0}
<div>
  <Historial bind:history={history}/>
</div>
{:else}
<p>Cargando historial...</p>
{/if}


