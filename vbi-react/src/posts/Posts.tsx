export default function Posts() {
  const posts = [
    {
      title: "Why I Started This Blog",
      summary:
        "A quick dive into my motivation for sharing what I learn as a dev.",
      date: "May 19, 2025",
      tags: ["Personal", "Journey"],
    },
    {
      title: "Lessons from Building My First API in Go",
      summary: "Things I learned (and wish I knew) when designing my backend.",
      date: "May 10, 2025",
      tags: ["Go", "Backend", "API"],
    },
  ];

  return (
    <div className="w-full min-h-screen w-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100 flex flex-col px-4 pt-24 pb-10">
      <main className="w-full max-w-6xl mx-auto flex-grow">
        <h1 className="text-4xl font-bold font-mono mb-8 border-b pb-2 dark:border-gray-700">
          Posts
        </h1>

        <div className="space-y-8 w-full">
          {posts.map((post, idx) => (
            <div
              key={idx}
              className="bg-white w-full dark:bg-gray-800 rounded-lg overflow-hidden shadow-sm"
            >
              <div className="p-6 w-full">
                <div className="flex flex-col justify-between items-start w-full sm:flex-row">
                  <h2 className="font-mono text-lg font-semibold sm:text-2xl">
                    {post.title}
                  </h2>
                  <p className="mb-3">{post.date}</p>
                </div>
                <p className="text-gray-700 dark:text-gray-300 mb-4 w-full">
                  {post.summary}
                </p>
                <div className="flex flex-wrap gap-2 mt-4">
                  {post.tags.map((tag, tIdx) => (
                    <span
                      key={tIdx}
                      className="px-3 py-1 bg-gray-100 dark:bg-gray-700 text-gray-300 text-xs font-medium rounded-full"
                    >
                      {tag}
                    </span>
                  ))}
                </div>
              </div>
            </div>
          ))}

          {posts.length === 0 && (
            <div className="text-center py-12">
              <p className="text-gray-500 dark:text-gray-400 text-lg">
                No posts yet. Check back soon!
              </p>
            </div>
          )}
        </div>
      </main>
    </div>
  );
}
