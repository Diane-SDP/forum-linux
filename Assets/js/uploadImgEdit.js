document.addEventListener('DOMContentLoaded', function() {
    const mediaInput = document.getElementById('mediaInput');
    const uploadedImage = document.getElementById('postImage');
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
  
    if (uploadedImage.src !== '.' && uploadedImage.src.includes('.')) {
        removeImageBtn.style.display = 'flex';
    } else {
        removeImageBtn.style.display = 'none';
    }
    
   
    removeImageBtn.addEventListener('click', function(e) {
        e.preventDefault();
        var path = window.location.href;
        var parts = path.split('/');
        var id = parts.pop();
        console.log(id)
        uploadedImage.src = ''; 
        uploadedImage.style.display = 'none'; 
        removeImageBtn.style.display = 'none'; 
        mediaInput.value = ''; 
         
        fetch('/deleteImage/' + id, {
            method: 'POST', 
            headers: {
                'Content-Type': 'application/json' 
            },
        })
        .then(response => response.text())
        .then(data => {
            console.log(data); 
        })
        .catch(error => {
            console.error('Erreur lors de la suppression de l\'image:', error);
        });
    });
  });
  