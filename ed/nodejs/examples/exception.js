
try {
    throw {code: 701, messame: 'AnError'};
} catch (e) {
    if (e.code = 702) {
        console.error('CAUGHT ERROR !!! :', e);
    }
    // throw e; // won't work
}

throw new Error('MyNewError');
