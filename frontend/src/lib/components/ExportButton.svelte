<script lang="ts">
 export let noteId:string;
 async function doExport(format:string){ /* call export API */ }
</script>

<button on:click={() => doExport("txt")}>Export TXT</button>
<button on:click={() => doExport("pdf")}>Export PDF</button>