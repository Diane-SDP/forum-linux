document.addEventListener('DOMContentLoaded', function() {
  const mediaInput = document.getElementById('mediaInput');
  const uploadedImage = document.getElementById('uploadedImage');
  const removeImageBtn = document.getElementById('removeImage');

  mediaInput.addEventListener('change', function(event) {
      const file = event.target.files[0];
      const reader = new FileReader();

      reader.onload = function(e) {
          uploadedImage.src = e.target.result;
          uploadedImage.style.display = 'block';
          removeImageBtn.style.display = 'flex'; 
      };
      reader.readAsDataURL(file);
  });

  removeImageBtn.addEventListener('click', function(e) {
    e.preventDefault();
      uploadedImage.src = ''; 
      uploadedImage.style.display = 'none'; 
      removeImageBtn.style.display = 'none'; 
      mediaInput.value = '';
  });
});
