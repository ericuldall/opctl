import React from 'react';
import Description from './Description';

export default ({
                  name,
                  param,
                  pkgRef,
                  value,
                }) => {
  return (
    <div className='form-group'>
      <label className='form-control-label' htmlFor={name}>{name}</label>
      <Description value={param.description} pkgRef={pkgRef}/>
      <input
        className='form-control'
        id={name}
        readOnly={true}
        type='text'
        value={value || param.default || ""}
      />
    </div>
  );
}
