$(function (global) {
    var search     = window.location.search;
    var parameters = {};

    search.substr(1).split('&').forEach(function (entry) {
        var eq = entry.indexOf('=');
        if (eq >= 0) {
            parameters[decodeURIComponent(entry.slice(0, eq))] =
                decodeURIComponent(entry.slice(eq + 1));
        }
    });

    if (parameters.variables) {
        try {
            parameters.variables =
                JSON.stringify(JSON.parse(query.variables), null, 2);
        } catch (e) {
            // Do nothing
        }
    }

    function onEditQuery(newQuery) {
        parameters.query = newQuery;
        updateURL();
    }

    function onEditVariables(newVariables) {
        parameters.variables = newVariables;
        updateURL();
    }

    function updateURL() {
        var newSearch = '?' + Object.keys(parameters).map(function (key) {
                return encodeURIComponent(key) + '=' +
                    encodeURIComponent(parameters[key]);
            }).join('&');
        history.replaceState(null, null, newSearch);
    }

    function graphQLFetcher(graphQLParams) {
        return fetch(window.location.origin + '/graphql', {
            method : 'post',
            headers: {'Content-Type': 'application/json'},
            body   : JSON.stringify(graphQLParams)
        }).then(function (response) {
            return response.json()
        });
    }

    global.renderGraphiql = function (elem) {
        var toolbar = React.createElement(GraphiQL.Toolbar, {}, [
            "Source available at ",
            React.createElement("a", {
                href: "https://github.com/rodlps22/todo-grapthql"
            }, "github")
        ]);
        React.render(
            React.createElement(GraphiQL, {
                fetcher        : graphQLFetcher,
                query          : parameters.query,
                variables      : parameters.variables,
                onEditQuery    : onEditQuery,
                onEditVariables: onEditVariables,
                defaultQuery   : "{\n"+
                    "  todoList(lang: \"GO\") {\n"+
                    "    id,\n"+
                    "    text,\n"+
                    "    done,\n"+
                    "    language {\n"+
                    "      name\n"+
                    "    }\n"+
                    "  },\n"+
                    "  todo(id: \"57c0044362f31f04777f743f\") {\n"+
                    "    id,\n"+
                    "    text,\n"+
                    "    done\n"+
                    "  }\n"+
                    "}\n"
            }, toolbar),
            elem
        );
    }
}(window));