<template>
  <div id="app">
    <!-- formulario de busqueda -->
    <input v-model="searchTerm" placeholder="Ingrese el término de búsqueda" />
    <button @click="search">Buscar</button>

    <!-- Tabla -->
    <!-- <table class="email-table" v-if="emails.length"> -->
      <table>
  <thead>
    <tr>
      <th>Email</th>
    </tr>
  </thead>
  <tbody>
    <!-- Recorre emails para mostrar tablas -->
    <template v-for="(source, index) in sources" :key="index">
      <tr v-if="source.content">
        <td>{{ source.content }}</td>
      </tr>
    </template>
  </tbody>
</table>
    <!-- <div v-else>
      <p>No hay resultados para mostrar.</p>
    </div> -->
  </div>
</template>

<script lang="ts">
export default {
  data() {
    return {
      searchTerm: '',
      // Asignar el tipo a emails para que no lo tome como un valor nil
      sources: [] as { content: string }[],
    };
  },
  methods: {
    async search() {
      try {

      console.log(this.searchTerm);

      // Construye la URL con el término de búsqueda
       const url2 = `http://localhost:8080/search/${this.searchTerm}`;  
       // Codifica el termino para que sea aceptable por la url
       const url = `http://localhost:8080/search/${encodeURIComponent(this.searchTerm)}`;

       console.log(url2);
       console.log(url);
    
      // GET con fetch
      const response = await fetch(url)
          // Codigo de exito (código de estado 200)
          if (!response.ok) {
            throw new Error(`Error al realizar la solicitud: ${response.status}`);
          }
          // Respuesta a JSON
          const data = await response.json();

          //console.log(data)
          console.log(data.hits.hits._source)

          const ema = data.hits.hits.map(hit => hit._source);
          console.log(ema);

          // Actualiza la lista
          this.sources = ema;
          //console.log(this.emails)
        } catch (error: any) {
          console.error(error.message);
    }
  },
}
}
</script>
