$(function () {
  function initAjaxForm() {
    $('body').on('submit', '.ajaxForm', function (e) {
      e.preventDefault();
      $.ajax({
        type: $(this).attr('method'),
        url: $(this).attr('action'),
        data: $(this).serialize()
      })
      .done(function (data) {
        if (typeof data.message !== 'undefined') {
          $('.ajaxForm').html(data.message);
        }
      })
      .fail(function (jqXHR, textStatus, errorThrown) {
        if (typeof jqXHR.responseJSON !== 'undefined') {
          if (jqXHR.responseJSON.hasOwnProperty('form')) {
            $('.formBody').html(jqXHR.responseJSON.form);
          }
          $('.formError').html(jqXHR.responseJSON.message);
        } else {
          $('.formErrors').html(errorThrown);
        }
      });
    });
  }
  initAjaxForm();
});
