document.querySelectorAll('.coming-soon-wrapper').forEach(wrapper => {
  const button = wrapper.querySelector('.coming-soon');
  button.addEventListener('click', e => {
    e.preventDefault();
    wrapper.classList.add('show-tooltip');
    clearTimeout(wrapper._tooltipTimeout);
    wrapper._tooltipTimeout = setTimeout(() => {
      wrapper.classList.remove('show-tooltip');
    }, 2000);
  });
});