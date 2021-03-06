
from steps.bdd_test_util import cli_call

def after_scenario(context, scenario):
	if 'doNotDecompose' in scenario.tags:
		print("Not going to decompose after scenario {0}, with yaml '{1}'".format(scenario.name, context.compose_yaml))
	else:
		if 'compose_yaml' in context:
			print("Decomposing with yaml '{0}' after scenario {1}, ".format(context.compose_yaml, scenario.name))
			context.compose_output, context.compose_error, context.compose_returncode = \
        		cli_call(context, ["docker-compose", "-f", context.compose_yaml, "kill"], expect_success=True)
			context.compose_output, context.compose_error, context.compose_returncode = \
        		cli_call(context, ["docker-compose", "-f", context.compose_yaml, "rm","-f"], expect_success=True)


