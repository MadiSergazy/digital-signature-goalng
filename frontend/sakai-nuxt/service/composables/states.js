export const states = {
    jwt_token() {
        return useState('jwt_token');
    },
    setJWT(new_jwt_token) {
        return useState('jwt_token', () => new_jwt_token);
    },
    get_email() {
        return useState('email');
    },
    set_email(new_email) {
        return useState('email', () => new_email);
    },
    get_window_width() {
        return useState('window_width');
    },
    set_window_width(new_window_width) {
        return useState('window_width', () => new_window_width);
    }
};
