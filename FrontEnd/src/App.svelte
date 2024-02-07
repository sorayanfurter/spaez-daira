<script context="module" lang="ts">
  import type { AxiosResponse } from "axios";
  import { onMount } from "svelte";
  import type { HistoryEntry } from "./Componentes/Historial.svelte";
  import Display from "./Componentes/Display.svelte";
  import BotonNumerico from "./Componentes/BotonNumerico.svelte";
  import BotonOperacion from "./Componentes/BotonOperacion.svelte";
  import Historial from "./Componentes/Historial.svelte";
  import axios from "axios";
</script>

<script lang="ts">
  let input: string = "";
  let result: number = 0;
  let history: HistoryEntry[] = [];
  $: {
    input;
    result;
    history;
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

      console.log("Response data:", response.data);
      history = response.data;
    } catch (error) {
      console.error("Error fetching history:", error);
    }
  }

  onMount(() => {
    fetchHistory();
  });

  async function handleButtonClick(event: {
    detail: string | number;
  }): Promise<void> {
    const clickedValue = event.detail;

    if (clickedValue === "=") {
      if (input.length > 0) {
        try {
          console.log("Fetching history...");
          const response: AxiosResponse<{ resultado: number }> =
            await axios.post("http://localhost:9090/Calcular", {
              n1: parseFloat(input.split(" ")[0]),
              n2: parseFloat(input.split(" ")[2]),
              operacion: input.split(" ")[1],
            });
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

<div class="grid grid-cols-2 gap-4 p-4">
  <div
    class="grid grid-rows-1 gap-4 bg-white border-2 border-black justify-center m-20 p-8"
  >
    <div
      class="bg-gray-200 text-black font-bold text-center text-3xl"
    >
      Calculadora Basica
    </div>
    <div class="grid grid-rows-1 gap-4 bg-white justify-center">
      <!-- Display -->
      <div class="bg-black text-white p-2 text-right">
        <Display bind:input bind:result />
      </div>

      <div class="grid grid-cols-4 grid-rows-5">
        <!-- Fila 1 -->
        <div class="w-full p-1">
          <BotonOperacion operation="CE" on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="C" on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="+/-" on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="%" on:buttonClick={handleButtonClick} />
        </div>

        <!-- Fila 2 -->

        <div class="w-full p-1">
          <BotonNumerico value={7} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonNumerico value={8} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonNumerico value={9} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="/" on:buttonClick={handleButtonClick} />
        </div>

        <!-- Fila 3 -->
        <div class="w-full p-1">
          <BotonNumerico value={4} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonNumerico value={5} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonNumerico value={6} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="*" on:buttonClick={handleButtonClick} />
        </div>

        <!-- Fila 4 -->

        <div class="w-full p-1">
          <BotonNumerico value={1} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonNumerico value={2} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonNumerico value={3} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="-" on:buttonClick={handleButtonClick} />
        </div>

        <!-- Fila 5 -->
        <div class="w-full p-1">
          <BotonNumerico value={0} on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="." on:buttonClick={handleButtonClick}/>
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="=" on:buttonClick={handleButtonClick} />
        </div>
        <div class="w-full p-1">
          <BotonOperacion operation="+" on:buttonClick={handleButtonClick} />
        </div>
      </div>
    </div>
  </div>
  <div class="col-span-1 p-10">
    <!-- Historial -->
    {#if history && history.length > 0}
      <div class="bg-white border-2 border-black p-10">
        <div
          class="bg-gray-200 p-10 text-black font-bold text-center text-3xl"
        >
          Calculadora Basica
        </div>
        <h2 class=" font-bold mb-4 text-center text-3xl pt-8">Historial</h2>
        <table class="min-w-full bg-black text-white">
          <thead>
            <tr>
              <th class="border-b-2 border-gray-400 py-2">Fecha</th>
              <th class="border-b-2 border-gray-400 py-2">Operación</th>
              <th class="border-b-2 border-gray-400 py-2">Resultado</th>
            </tr>
          </thead>
          <tbody>
            {#each history as entry}
              <tr class="border-b border-gray-300">
                <td class="py-2">{entry.fecha}</td>
                <td class="py-2">{entry.operacion}</td>
                <td class="py-2">{entry.resultado}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {:else}
      <p>Cargando historial...</p>
    {/if}
  </div>
</div>

<style>
  @import "tailwindcss/base.css";
  @import "tailwindcss/components.css";
  @import "tailwindcss/utilities.css";
</style>
