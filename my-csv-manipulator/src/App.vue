<template>
  <div class="container">
    <h1 class="text-center my-4">CSV Manipulator</h1>

    <div class="mb-4">
      <input type="file" ref="fileInput" @change="uploadFile" />
      <button class="btn btn-primary" @click="sortCSV">Sort CSV</button>
      <button class="btn btn-success" @click="downloadCSV">Download Sorted CSV</button>
    </div>

    <table class="table table-striped" v-if="csvPreview.length">
      <thead>
      <tr>
        <th v-for="(col, index) in csvPreview[0]" :key="index">{{ col }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(row, index) in csvPreview.slice(1)" :key="index">
        <td v-for="(cell, index) in row" :key="index">{{ cell }}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      file: null,
      csvPreview: [],
      downloadLink: "",
      app_error: "",
    }
  },
  methods: {
    uploadFile() {
      this.file = this.$refs.fileInput.files[0]
      const reader = new FileReader();
      reader.readAsText(this.file, 'UTF-8');
      reader.onload = (evt) => {
        const text = evt.target.result;
        this.csvPreview = text.split('\n').map(row => row.split(','));
      }
    },
    async makeRequest(endpoint, columnToSort) {
      if (!this.file) {
        this.app_error = "No file uploaded.";
        return;
      }

      this.app_error = "";
      const formData = new FormData();
      formData.append('file', this.file);

      try {
        const response = await fetch(`http://localhost:8080/${endpoint}?column=` + columnToSort, {
          method: 'POST',
          body: formData,
        });
        if (response.ok) {
          const blob = await response.blob();
          this.downloadLink = URL.createObjectURL(blob);

          const text = await blob.text();
          this.csvPreview = text.split('\n').map(row => row.split(','));
        } else {
          this.app_error = "Something went wrong.";
        }
      } catch (e) {
        this.app_error = e.toString();
      }
    },
    async sortCSV() {
      await this.makeRequest("sort-csv", 0);
    },
    downloadCSV() {
      if (!this.downloadLink) {
        this.app_error = "No file to download.";
        return;
      }
      const link = document.createElement('a');
      link.href = this.downloadLink;
      link.download = "processed.csv"
      link.click();
    }
  }
}
</script>

<style>

</style>
