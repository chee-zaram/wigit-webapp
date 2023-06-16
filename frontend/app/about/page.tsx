// about page

export const metadata = { title: 'wigit About' };

const AboutUs = () => (
    <div className="about_page w-full min-h-screen p-6 md:p-12">
        <div className="text-accent about_page_header w-full h-[40vh]">
            <h1 className="text-4xl font-extrabold">About Us</h1>
            <p>We are a wigging company based in Ashanti region, Ghana.</p>
        </div>
        <div className='w-[80vw] mx-auto rounded md:flex'>
            <div className="md:w-1/3 mb-8 md:mb-0 md:mr-8 min-h-[15vh] bg-dark_bg/40">
                <p>Lorem Ipsum</p>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam ac orci condimentum, luctus ligula vel, rutrum nisl. Ut tincidunt risus quis enim sodales auctor. Nam non diam dui. Nam scelerisque orci sit amet elit sodales tristique. Maecenas varius libero ex, quis porta nisl gravida nec. Aliquam erat volutpat. Aenean at tempor nisl. Duis sollicitudin dolor eu mauris finibus mollis. Donec sagittis tortor vel consectetur feugiat. Etiam vitae feugiat quam. Sed id velit id leo vestibulum ultricies. Duis id velit molestie, rutrum diam ut, tempus sem. Curabitur mollis dolor velit, non porta ante tempus id. Fusce sed leo aliquet, accumsan tellus id, posuere lorem.</p>
            </div>
            <div className="md:w-2/3 selection:min-h-[15vh] bg-accent/40">
                <p>Lorem Ipsum</p>
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam ac orci condimentum, luctus ligula vel, rutrum nisl. Ut tincidunt risus quis enim sodales auctor. Nam non diam dui. Nam scelerisque orci sit amet elit sodales tristique. Maecenas varius libero ex, quis porta nisl gravida nec. Aliquam erat volutpat. Aenean at tempor nisl. Duis sollicitudin dolor eu mauris finibus mollis. Donec sagittis tortor vel consectetur feugiat. Etiam vitae feugiat quam. Sed id velit id leo vestibulum ultricies. Duis id velit molestie, rutrum diam ut, tempus sem. Curabitur mollis dolor velit, non porta ante tempus id. Fusce sed leo aliquet, accumsan tellus id, posuere lorem.</p>
            </div>
            
        </div>
    </div>
)

export default AboutUs;
