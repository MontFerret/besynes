module.exports = {
    module: {
        rules: [
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader'],
            },
            {
                test: /\.module.css$/i,
                loader: 'css-loader',
                options: {
                    modules: true,
                },
            },
        ],
    },
};
