import { Router } from 'itty-router';

// Create a new router
const router = Router();

/*
Our index route, a simple hello world.
*/
router.get('/ads', () => {
	const data = [{
		ad:"zJunior where you can find you first intern",
		topics:['jobs','software engineering']
		},
		{
			ad:"onboardbase where security meets the collaboration",
			topics:['security','collaboration','frontend','devtool']
		}
	]
	return new Response(JSON.stringify(data),
		{
			headers: {
				"content-type": "application/json;charset=UTF-8",
		}}
	  );
});

router.all('*', () => new Response('404, not found!', { status: 404 }));

export default {
	fetch: router.handle,
};
