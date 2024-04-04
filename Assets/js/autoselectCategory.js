function getCategoryFromPreviousUrl() {
    const previousUrl = document.referrer;
    const matches = previousUrl.match(/\/topic\/(\d+)/);
    if (matches && matches.length > 1) {
        return parseInt(matches[1]); 
    }
    return null;
}

function selectCategoryOption(categoryId) {
    const selectElement = document.getElementById('post-category');
    if (!selectElement) return; 

    for (let option of selectElement.options) {
        if (parseInt(option.value) === categoryId) {
            option.selected = true; 
            break;
        }
    }
}

window.onload = function() {
    const categoryId = getCategoryFromPreviousUrl();
    if (categoryId !== null) {
        selectCategoryOption(categoryId);
    }
};
